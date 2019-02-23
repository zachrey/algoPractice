package main

import "fmt"

// 基数排列
func radixSort(arr []int) {
	n := len(arr)
	buckets := [10][]int{}
	max := arr[0]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	// 需要排列的次数
	digit := 0
	for max > 0 {
		max /= 10
		digit++
	}

	start := 1 // 用于每个值的各个位

	for i := 0; i < digit; i++ {
		start *= 10
		// 根据每一位分别存入数据
		for j := 0; j < n; j++ {
			curIndex := arr[j] % start / (start / 10) // 取当前位
			if buckets[curIndex] == nil {
				buckets[curIndex] = []int{}
			}
			buckets[curIndex] = append(buckets[curIndex], arr[j])
		}

		// 取出数据
		k := 0
		for j := 0; j < len(buckets); j++ {
			for jj := 0; jj < len(buckets[j]); jj++ {
				arr[k] = buckets[j][jj]
				k++
			}
		}

		buckets = [10][]int{}
	}
}

func main() {
	arr := []int{1, 10, 100, 1000, 98, 67, 3, 28, 67, 888, 777}
	radixSort(arr)
	fmt.Println(arr)
}
