package main

import (
	"fmt"
	"sort"
)

func findMax(arr []int32, bus []int32) int32 {
	var max int32
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}

// return index
func findItem(arr []int32, f int32) int {
	for i, item := range arr {
		if item == f {
			return i
		}
	}
	return -1
}

func contains(arr []int32, f int32) bool {
	for _, item := range arr {
		if item == f {
			return true
		}
	}
	return false
}

func getServerIndex(n int32, arrival []int32, burstTime []int32) []int32 {
	// Write your code here
	servers := []int{1, 2, 3}
	usingServers := []int32{}
	freeTime := [3]int32{-1, -1, -1}
	processedIndex := []int32{}
	var logs []int32
	max := findMax(arrival, burstTime)
	for i := 0; i <= int(max)+1; i++ {

		index := findItem(arrival, int32(i))
		for j, ft := range freeTime {
			if ft != -1 && int(ft) <= i {
				freeTime[j] = -1
				index := findItem(usingServers, int32(j+1))
				take := usingServers[index]
				servers = append(servers, int(take))
				usingServers = append(usingServers[:index], usingServers[index+1:]...)
				sort.Sort(sort.IntSlice(servers))
			}
		}
		if index == -1 || contains(processedIndex, int32(index)) {
			continue
		}
		if len(servers) == 0 {
			arrival[index]++
			continue
		}
		sort.Sort(sort.IntSlice(servers))
		take := servers[0]
		servers = servers[1:]
		usingServers = append(usingServers, int32(take))
		freeTime[take-1] = int32(i) + burstTime[index]
		processedIndex = append(processedIndex, int32(index))
		logs = append(logs, int32(take))
	}
	return logs
}

func main() {
	arr := []int32{2, 4, 1, 8, 9}
	bus := []int32{7, 9, 2, 4, 5}
	fmt.Print(getServerIndex(3, arr, bus))
	// 1 2 1 3 2
}
