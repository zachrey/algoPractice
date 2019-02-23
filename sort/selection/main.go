package main

import "fmt"

func insertSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		index := i
		for j := i; j < n; j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func main() {
	arr := []int{6, 5, 4, 1, 2, 3}
	insertSort(arr)
	fmt.Println(arr)
}
