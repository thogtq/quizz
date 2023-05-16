package solution

// you can also use imports, for example:
// import "fmt"
import "sort"

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
	for _, item := range A {
		if item < min && item > 1 {
			min = item
		}
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
