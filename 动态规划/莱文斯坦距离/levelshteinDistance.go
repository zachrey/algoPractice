package al42distance

import (
	"math"
)

func lwstDP(a, b []string, n, m int) int {
	states := make([][]int, n)
	for i := range states {
		states[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		if a[0] == b[j] {
			states[0][j] = j
		} else if j != 0 {
			states[0][j] = states[0][j-1] + 1
		} else {
			states[0][j] = 1
		}
	}

	for i := 0; i < n; i++ {
		// 我这里注释的初始化，与下面的相等，这个更好理解，第一个元素到底是 0  还是 1
		// if i == 0 {
		// 	if a[i] == b[0] {
		// 		states[i][0] = 0
		// 	} else {
		// 		states[i][0] = 1
		// 	}
		// } else {
		// 	states[i][0] = states[i-1][0] + 1
		// }

		if b[0] == a[i] {
			states[i][0] = i
		} else if i != 0 {
			states[i][0] = states[i-1][0] + 1
		} else {
			states[i][0] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i] != b[j] {
				states[i][j] = min(states[i][j-1]+1, states[i-1][j]+1, states[i-1][j-1]+1)
			} else {
				states[i][j] = min(states[i][j-1]+1, states[i-1][j]+1, states[i-1][j-1])
			}
		}
	}

	return states[n-1][m-1]
}

func min(x, y, z int) int {
	min := math.MaxInt16
	if x < min {
		min = x
	}

	if y < min {
		min = y
	}

	if z < min {
		min = z
	}

	return min
}
