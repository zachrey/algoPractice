package al40triangle

import (
	"math"
)

func getMinValue(items [][]int, n int) int {
	// 初始化 备忘录
	states := make([][]int, n)
	for i := range states {
		states[i] = make([]int, i+1)
	}

	for i := range states {
		for j := range states[i] {
			states[i][j] = math.MaxInt16
		}
	}

	// 初始化第一层的状态
	states[0][0] = items[0][0]

	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ { // 遍历每一层
			// 取左上角的值
			if j-1 >= 0 {
				states[i][j] = states[i-1][j-1] + items[i][j]
			}
			// 取正上方的值
			if j < i {
				tv := states[i-1][j] + items[i][j] // 比较左上角和右上角的值谁更小
				if tv < states[i][j] {
					states[i][j] = tv
				}
			}
		}
	}

	// 获取最底层的最小值就是结果
	minV := math.MaxInt16
	for i := 0; i < n; i++ {
		if states[n-1][i] < minV {
			minV = states[n-1][i]
		}
	}

	return minV
}
