package main

import (
	"bufio"
	"os"
)

func main() {
	fileIn, _ := os.Open("input-201.txt")
	defer fileIn.Close()
	fscanner := bufio.NewScanner(fileIn)
	count := make(map[string]int)
	for fscanner.Scan() {
		count[fscanner.Text()]++
	}

	fileOut, _ := os.OpenFile("input-201.a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer fileOut.Close()
	result := ""
	for key, val := range count {
		if val < 2 {
			result += key + "\n"
		}
	}
	fileOut.WriteString(result)

}
