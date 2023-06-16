package main

import "strings"

func main() {
	println(Solution(">>>><.<"))
}

func scanLeftPassed(index int, S string) int {
	sub := S[0:index]
	return strings.Count(sub, ">")
}

func scanRightPassed(index int, S string) int {
	sub := S[index:]
	return strings.Count(sub, "<")
}

func Solution(S string) int {
	passes := 0
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "." {
			passes += scanLeftPassed(i, S)
			passes += scanRightPassed(i, S)
		}
	}
	return passes
}
