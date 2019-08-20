package main

import (
	"fmt"
	"strconv"
)

func main() {
	test()
}

func reverse(x int) int {
	var (
		result int
		change = false
		locate int
	)
	s := strconv.Itoa(x)
	str := []byte(s)
	lens := len(str)
	if string(str[0]) == "-" {
		str = str[1:lens]
		change = true
		lens = lens - 1
	}
	for i := 0; i < lens/2; i++ {
		a := str[i]
		str[i] = str[lens-1-i]
		str[lens-1-i] = a
	}

	for i := 1; i < lens; i++ {
		if str[i] == 0 {
			locate = i + 1
			continue
		}

		break
	}

	str = str[locate:lens]
	if change {
		str = append([]byte("-"), str...)
	}
	result, err := strconv.Atoi(string(str))
	if err != nil || result > 1<<31-1 || result < -(1<<31) {
		return 0
	}
	return result
}

func test() {
	// fmt.Print(strconv.Itoa(-100))
	n := 1<<31 - 1
	fmt.Print(n)
}
