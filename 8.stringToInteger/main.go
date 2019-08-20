package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(myAtoi("20000000000000000000"))
}

func myAtoi(str string) int {
	var (
		reStr     []byte
		condition = false
		hadBlank  = false
	)
	s := []byte(str)

HERE:
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case " ":
			if hadBlank {
				break HERE
			}
			continue
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			reStr = append(reStr, s[i])
			if !condition {
				condition = true
			}
			if !hadBlank {
				hadBlank = true
			}
		case "-", "+":
			if !hadBlank {
				hadBlank = true
			}
			if condition {
				break HERE
			}
			reStr = append(reStr, s[i])
			condition = true
		default:
			break HERE
		}
	}

	r, err := strconv.Atoi(string(reStr))
	if err != nil {
		// fmt.Print(err, "\n")
		// fmt.Print(strings.Split(err.Error(), ":")[2], "\n")
		if strings.Split(err.Error(), ":")[2] == " value out of range" {
			if string(reStr[0]) == "-" {
				return -(1 << 31)
			}
			return 1<<31 - 1

		}
		return 0
	}

	if r > 1<<31-1 {
		return 1<<31 - 1
	} else if r < -(1 << 31) {
		return -(1 << 31)
	}
	return r
}
