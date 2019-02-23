package main

import "fmt"

func insertSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}

func main() {
	arr := []int{6, 5, 4, 1, 2, 3}
	insertSort(arr)
	fmt.Println(arr)
}
