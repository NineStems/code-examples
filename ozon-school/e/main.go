package main

import (
	"fmt"
	"sync"
)

type dataCh struct {
	x   int
	y   int
	okX bool
	okY bool
}

func main() {
	n := 5
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)

	//Merge2Channels(foo, in1, in2, out, n)

	go Merge2Channels(foo, in1, in2, out, n)
	in1 <- 1
	in1 <- 2
	in2 <- 1
	in2 <- 2
	in1 <- 3
	in2 <- 3
	in2 <- 4
	in2 <- 5
	in1 <- 4
	in1 <- 5
	for i := 0; i < n; i++ {
		fmt.Println(<-out)
	}
}

func foo(i int) int {
	return i
}

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	readChannels(f, in1, in2, n, out, &wg)
	wg.Wait()
}

func merge2Channels(f func(int) int, in1 int, in2 int, out chan<- int) {
	out <- f(in1) + f(in2)
}

func readChannels(f func(int) int, in1 <-chan int, in2 <-chan int, n int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu := sync.Mutex{}
	mu.Lock()
	data := []dataCh{}
	for i := 0; i < n; i++ {
		data = append(data, dataCh{})
	}
	mu.Unlock()
	countX := 0
	countY := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if countX < n {
				select {
				case value1, _ := <-in1:
					mu.Lock()
					data[countX].x = value1
					data[countX].okX = true
					countX++
					mu.Unlock()
				}
			} else {
				break
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if countY < n {
				select {
				case value2, _ := <-in2:
					mu.Lock()
					data[countY].y = value2
					data[countY].okY = true
					countY++
					mu.Unlock()
				}
			} else {
				break
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := n
		for {
			if counter == 0 {
				break
			}
			mu.Lock()
			for idx, val := range data {
				if val.okX && val.okY {
					go merge2Channels(f, val.x, val.y, out)
					data[idx].okX = false
					data[idx].okY = false
					counter--
					break
				}
			}
			mu.Unlock()
		}
	}()
}
