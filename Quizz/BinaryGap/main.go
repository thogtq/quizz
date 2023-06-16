package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print(Solution(295937))
	// 100100001000000000 gap = 4
	// [ 00 0000 000000000]
}

func findMaxGap(arr []string) int {
	max := 0
	for i, item := range arr {
		fmt.Println(item)
		if len(item) > max {
			if i == len(arr)-1 {
				continue
			}
			max = len(item)
		}
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
