package main

import "fmt"

/*
基础的二分查找

数列的前提条件
1. 有序
2. 不易过大，比如1G大的数据，因为底层使用数组实现，所以需要连续的内存才行
3. 不易过小，过小的话使用遍历查找就行，没必要做很多次的计算，出发减少比较能很见效的减少时间

时间复杂度：O(logn)
*/

func binarySearch(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}

	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > val {
			right = mid - 1
		} else if arr[mid] < val {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	index := binarySearch(arr, 3)
	fmt.Println(index)
}
