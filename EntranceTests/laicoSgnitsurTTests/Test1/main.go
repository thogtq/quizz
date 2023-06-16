package main

import "fmt"

func reverseArr(arr []int32, start int, end int) {
	e := end
	s := start
	for i := 0; i <= (end-start)/2; i++ {
		temp := arr[s]
		arr[s] = arr[e]
		arr[e] = temp
		s++
		e--
	}
}

func performOperations(arr []int32, operations [][]int32) []int32 {
	for i := 0; i < len(operations); i++ {
		start := operations[i][0]
		end := operations[i][1]
		reverseArr(arr, int(start), int(end))
	}
	return arr
}

func main() {
	arr := []int32{1, 2, 3}
	ope := [][]int32{{0, 2}, {1, 2}, {0, 2}}
	fmt.Println(performOperations(arr, ope))
}
