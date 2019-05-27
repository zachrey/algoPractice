package al42maxlen

import (
	"math"
)

func getMaxLen(a, b []string, n, m int) int {
	states := make([][]int, n)
	for i := range states {
		states[i] = make([]int, m)
	}

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		states[i][j] = -1
	// 	}
	// }

	for i := 0; i < m; i++ {
		if b[i] == a[0] {
			states[0][i] = 1
		} else if i > 0 {
			states[0][i] = states[0][i-1]
		} else {
			states[0][i] = 0
		}
	}

	for i := 0; i < n; i++ {
		if a[i] == b[0] {
			states[i][0] = 1
		} else if i > 0 {
			states[i][0] = states[i-1][0]
		} else {
			states[i][0] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i] == b[j] {
				states[i][j] = max(states[i-1][j], states[i][j-1], states[i-1][j-1]+1)
			} else {
				states[i][j] = max(states[i-1][j], states[i][j-1], states[i-1][j-1])
			}
		}
	}

	return states[n-1][m-1]
}

func max(x, y, z int) int {
	max := math.MinInt16
	if x > max {
		max = x
	}

	if y > max {
		max = y
	}

	if z > max {
		max = z
	}

	return max
}
