package main

import "fmt"

func bubbleSort(arr []int) {
	for i, n := 0, len(arr); i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if flag {
			break
		}
	}
}

func main() {
	arr := []int{4, 5, 6, 3, 2, 1}
	bubbleSort(arr)
	fmt.Println(arr)
}
