package al42lwst

import (
	"math"
)

var edist = math.MaxInt16

func lwst(i, j, n, m, cv int, a, b []string) {
	if i == n || j == m {
		if i < n {
			cv += (n - i)
		}
		if j < m {
			cv += (m - j)
		}
		if cv < edist {
			edist = cv
		}
		return
	}

	if a[i] == b[j] {
		lwst(i+1, j+1, n, m, cv, a, b) // 相等的一种情况
	} else {
		lwst(i+1, j, n, m, cv+1, a, b) // 不相等的三种情况
		lwst(i, j+1, n, m, cv+1, a, b)
		lwst(i+1, j+1, n, m, cv+1, a, b)
	}
}
