package main

import "fmt"

/**
计数排序
*/

func countingSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	max := arr[0]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	count := make([]int, max+1)
	// 每个下标对应的是arr数组里面的值，count的元素是每个arr元素的个数
	for i := 0; i < n; i++ {
		count[arr[i]]++
	}

	// 顺序累加
	for i := 1; i < max+1; i++ {
		count[i] = count[i-1] + count[i]
	}

	temp := make([]int, n)
	for i := 0; i < n; i++ {
		index := arr[i] - 1 // 注意这需要减1
		temp[count[index]] = arr[i]
		count[arr[i]]--
	}

	for i := 0; i < n; i++ {
		arr[i] = temp[i]
	}
}

func main() {
	arr := []int{2, 5, 3, 0, 2, 3, 0, 3}
	countingSort(arr)
	fmt.Println(arr)
}
