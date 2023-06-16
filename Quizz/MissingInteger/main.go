package main

// you can also use imports, for example:
// import "fmt"
import (
	"fmt"
	"sort"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func isInArr(A []int, value int, current int) bool {
	for i := current; i < len(A); i++ {
		if A[i] == value {
			return true
		}
	}
	return false
}
func Solution(A []int) int {
	min := 1
	sort.Sort(sort.IntSlice(A))
	if A[0] > 1 {
		min = A[0]
	}
	current := 0
	for {
		if isInArr(A, min, current) {
			min++
			current++
		} else {
			return min
		}
	}
}

func main() {
	arr := []int{-1, -3, -2} //1
	fmt.Println(Solution(arr))
	arr2 := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(Solution(arr2)) //5
}
