package al42maxlen

func getMaxLen(a []int, n int) int {
	states := make([]int, n)

	states[0] = 1

	for i := 1; i < n; i++ {
		max := 1
		for j := 0; j < i; j++ {
			if a[i] > a[j] && states[j] >= 1 {
				if states[j]+1 > max {
					max = states[j] + 1 // 取所以小于当前值的序列，取其中的最大值
				}
			}
		}
		states[i] = max // 当前下标的最长序列由前面所以值小于当前值的节点的最大的那个（子序列+1）
	}

	return states[n-1]
}
