package main

import (
	"fmt"
	s "strings"
)

func main() {
	examples1 := []string{
		"()",
		"({[]})",
		"((((",
		"[{]}",
		//")",
	}

	examples2 := []string{
		"()",
		"({[]})",
		"((((",
		"[{]}",
		")",
	}

	//need to find problem in work
	fmt.Println("first version from june-aug of 2020, its doesnot work at 100%")
	for _, ex := range examples1 {
		fmt.Println(ex, validate_v1(ex))
	}

	fmt.Println("second version from nov of 2020, example if using stack")
	for _, ex := range examples2 {
		fmt.Println(ex, validate_v2(ex))
	}
}

func validate_v1(str string) bool {
	strOpen := "({["
	strClose := "]})"
	vals := []string{"()", "[]", "{}"}
	lenMas := len(str)
	st := make([]string, lenMas)
	var counter int
	for _, val := range str {
		strVal := string(val)
		strTypeOpen := s.Index(strOpen, strVal)
		strTypeClose := s.Index(strClose, strVal)

		if strTypeOpen >= 0 {
			st[counter] = strVal
			counter++
		} else if strTypeClose >= 0 {
			counter--
			tempVal := st[counter]

			for _, valCh := range vals {

				if s.Index(valCh, tempVal) >= 0 && s.Index(valCh, strVal) >= 0 {
					break
				} else if (s.Index(valCh, tempVal) >= 0 && s.Index(valCh, strVal) < 0) || (s.Index(valCh, tempVal) < 0 && s.Index(valCh, strVal) >= 0) {
					return false
				} else {
					continue
				}
			}
			st[counter] = ""
		} else {
			return false
		}
	}
	return true

}

func validate_v2(s string) bool {
	data := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	stack := []string{}
	for _, val := range s {
		vs := string(val)
		v, ok := data[vs]
		if !ok {
			stack = append(stack, vs)
			continue
		}

		if v != "" && len(stack) == 0 {
			return false
		}

		if v != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1]

	}

	return len(stack) == 0
}