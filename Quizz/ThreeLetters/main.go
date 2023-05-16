// 66/100 NEED IMPROVING
package main

import "fmt"

func main() {
	fmt.Println(Solution(2, 2))
}

func Solution(A int, B int) string {
	result := ""
	limitA := 0
	limitB := 0
	for A > 0 || B > 0 {
		putA := A > 0 && B-A < 2
		putB := B > 0 && B-A >= 2
		// fmt.Println(putA, putB, A, B, limitA, limitB, result)
		if putA || B == 0 && A > 0 || limitB == 2 {
			result += "a"
			A--
			limitA++
			limitB = 0
		}
		if putB || A == 0 && B > 0 || limitA == 2 {
			result += "b"
			limitB++
			B--
			limitA = 0
		}
	}
	return result
}
