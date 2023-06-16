package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("isLAst", isLast5(0, "5005"))
	println(Solution(-500))
}

func fiveCount(str string) int {
	return strings.Count(str, "5")
}

func isLast5(i int, str string) bool {
	sub := str[i:]
	return strings.Count(sub, "5") < 2
}

func Solution(N int) int {
	negative := false
	if N < 0 {
		negative = true
	}
	str := strconv.Itoa(N)
	result := ""
	isRemoved := false
	for i := 0; i < len(str); i++ {
		skip := false
		digit, _ := strconv.Atoi(string(str[i]))

		if i < len(str)-1 && digit == 5 && !isRemoved {
			lastFive := isLast5(i, str)
			digitplus1, _ := strconv.Atoi(string(str[i+1]))

			if digitplus1 > digit && !negative || negative && digitplus1 < digit {
				isRemoved = true
				continue
			} else if !lastFive {
				skip = true
			}
		}
		if digit == 5 && !isRemoved && !skip {
			if !negative && fiveCount(str) > 1 {
				isRemoved = true
				// println(d)
				continue
			}
		}
		result += strconv.Itoa(digit)
	}
	if negative {
		result = "-" + result
	}
	iResult, _ := strconv.Atoi(result)
	return iResult
}
