package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	if i := bytes.IndexByte(data, ' '); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func DoIt() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	sc := bufio.NewScanner(file)
	sc.Split(ScanLines)
	var firstValue int
	mapValue := make(map[int]struct{})
	result := false
	for sc.Scan() {
		val, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		if firstValue == 0 {
			firstValue = val
			continue
		}
		_, ok := mapValue[firstValue-val]
		if ok {
			result = true
			break
		}
		mapValue[val] = struct{}{}
	}
	if result {
		_ = ioutil.WriteFile("output.txt", []byte("1"), 0666)
	} else {
		_ = ioutil.WriteFile("output.txt", []byte("0"), 0666)
	}
}

func main() {
	DoIt()
}
