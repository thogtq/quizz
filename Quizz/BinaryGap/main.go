package main

import (
	"fmt"
	"strings"
)

func main() {
	test := fmt.Sprintf("%b", 1376796946)
	fmt.Println(test)
	testArr := strings.Split(test, "1")
	fmt.Println(len(testArr))
	fmt.Print(findMaxGap(testArr))
}

func findMaxGap(arr []string) int {
	max := 0
	for i, item := range arr {
		if len(item) > max {
			if i == len(arr)-1 {
				continue
			}
			max = len(item)
		}
		fmt.Println("item", item)
	}
	return max
}

func Solution(N int) int {
	bN := fmt.Sprintf("%b", N)
	if strings.Count(bN, "1") == 1 {
		return 0
	}
	bNArr := strings.Split(bN, "1")
	return findMaxGap(bNArr)
}
