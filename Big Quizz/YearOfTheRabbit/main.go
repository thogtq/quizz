// 100/100
package main

func main() {
	a := []int{1, 1, 1, 1, 1}
	b := []int{1, 1, 1, 1, 1}
	println(Solution(a, b))
}

func rotate(arr []int) {
	l := len(arr)
	for i := 0; i < len(arr); i++ {
		arr[i]--
		if arr[i] == -1 {
			arr[i] = l - 1
		}
	}
	// fmt.Println(arr)
}

func Solution(A []int, B []int) int {
	indexArr := []int{}
	rotateTimes := 0
	for i := 0; i < len(A); i++ {
		indexArr = append(indexArr, i)
	}

	for i := 0; i < len(A); {
		if A[i] == B[indexArr[i]] {
			// fmt.Println(A[i], B[indexArr[i]], rotateTimes)
			rotate(indexArr)
			i = 0
			rotateTimes++
			if rotateTimes > len(A)-1 {
				return -1
			}
			continue
		}
		i++
	}
	return rotateTimes
}
