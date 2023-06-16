package main

import "fmt"

func main() {
	input := []int64{5, 6, 2}
	fmt.Print(taskOfPairing(input))
}

func taskOfPairing(freq []int64) int64 {
	var pairs int64
	for i := 0; i < len(freq); i++ {
		pairs += freq[i] / 2
		left := freq[i] % 2
		freq[i] = left
		if i != len(freq)-1 && freq[i] == 1 && freq[i+1] > 1 {
			pairs += 1
			freq[i]--
			freq[i+1]--
		}
	}
	return pairs
}
