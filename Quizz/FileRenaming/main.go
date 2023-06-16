package main

import "fmt"

func main() {
	println(renameFile("abc", "aaabbbccc"))
	// println(renameFile("ccc", "cccc"))
}

func contains(arr []int, i int) bool {
	for _, item := range arr {
		if item == i {
			return true
		}
	}
	return false
}

func renameFile(newName string, oldName string) int32 {
	count := 0
	// offset := len(oldName) - len(newName)

	pickIndex := []int{}
	currentIndex := len(newName) - 1
	for i := 0; i < len(newName); i++ {
		pickIndex = append(pickIndex, i)
	}
	for {
		fmt.Println(pickIndex)
		word := ""
		for _, item := range pickIndex {
			word += string(oldName[item])
		}
		fmt.Println(word)
		if word == newName {
			count++
		}
		for {
			newVal := pickIndex[currentIndex] + 1
			if !contains(pickIndex, newVal) {
				pickIndex[currentIndex]++
				break
			}
			pickIndex[currentIndex]++
		}

		if pickIndex[currentIndex] == len(oldName) {
			if currentIndex == 0 {
				break
			}
			pickIndex[currentIndex] = currentIndex
			currentIndex--
			for {
				newVal := pickIndex[currentIndex] + 1
				if !contains(pickIndex, newVal) {
					pickIndex[currentIndex]++
					break
				}
				pickIndex[currentIndex]++
			}
		}
	}
	return int32(count)
}
