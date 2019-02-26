package main

import "fmt"

/*
二分查找变种一：
在有序数列里面查找第一个值为n的数
*/
func binarySearchFirst(arr []int, val int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != val { // 如果当前找到的值位于第一个或者前一个值不是该数值，表面该值就是第一个
				return mid
			}
			high = mid - 1
		}

	}
	return -1
}

/*
二分查找变种2：
查找数列中最后一个值为val的下标
*/
func binarySearchLast(arr []int, val int) int {
	n := len(arr)
	low, high := 0, n
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			if mid == n-1 || arr[mid+1] != val {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

/*
二分查找变种3：
查找数列中第一个大于等于 val 的下标
*/
func binarySearchFirst1(arr []int, val int) int {
	n := len(arr)
	low, high := 0, n
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < val {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

/*
二分查找变种4：
查找数列中第一个小于等于 val 的下标
...
*/

func main() {
	arr := []int{2, 2, 2, 3, 4, 5, 6, 8, 9, 10}
	// index := binarySearchFirst(arr, 2)
	// index := binarySearchLast(arr, 2)
	index := binarySearchFirst1(arr, 7)
	fmt.Println(index)
}
