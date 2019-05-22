package al40pack

/*
weight 每个物品的重量
n 物品的个数
w 背包的总重量
*/
func knapsack(weight []int, n, w int) int {
	// 声明一个二维数组保存状态
	states := make([][]bool, n)
	for i := range states {
		states[i] = make([]bool, w+1)
	}

	states[0][0] = true         // 第一个元素不装入
	states[0][weight[0]] = true // 第二个元素装入

	for i := 1; i < n; i++ { // 没一层的状态

		for j := 0; j <= w; j++ { // 不把第i个元素装入
			if states[i-1][j] == true {
				states[i][j] = states[i-1][j] // 这一层的重量不发生变化，所以还是上一层那么重
			}
		}

		for j := 0; j <= w-weight[i]; j++ { // 把第i个元素装入
			if states[i-1][j] == true {
				states[i][j+weight[i]] = true // 根据上一层的重量添加当前元素的重量，对应到指定的下标
			}
		}
	}

	// 最后一层的最右边被记录为true就是最大值
	for i := w; i >= 0; i-- {
		if states[n-1][i] == true {
			return i
		}
	}
	return 0
}

func knapsack2(weight []int, n, w int) int {
	states := make([]bool, w+1)

	states[0] = true
	states[weight[0]] = true

	for i := 1; i < n; i++ {

		for j := w - weight[i]; j >= 0; j-- {
			if states[j] == true {
				states[j+weight[i]] = true
			}
		}
	}

	// 最后一层的最右边被记录为true就是最大值
	for i := w; i >= 0; i-- {
		if states[i] == true {
			return i
		}
	}
	return 0
}

var (
	items  = []int{2, 2, 4, 6, 3} // 每个物品的重量
	values = []int{3, 4, 8, 9, 6} // 每个物品的价值
	n, w   = 5, 9                 // 总共5个物品，背包容量为9
)

func knapsack3() int {
	states := make([][]int, n)
	for i := range states {
		states[i] = make([]int, w+1)
		for j := range states[i] {
			states[i][j] = -1
		}
	}

	states[0][0] = 0
	states[0][items[0]] = values[0]

	for i := 1; i < n; i++ {

		for j := 0; j <= w; j++ { // 不装入第i个物品
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= w-items[i]; j++ { // 装入第i个物品

			if states[i-1][j] >= 0 {
				tv := states[i-1][j] + values[i]
				if tv > states[i][j+items[i]] { // 如果
					states[i][j+items[i]] = tv
				}
			}
		}
	}

	maxvalue := -1
	for j := 0; j <= w; j++ {
		if states[n-1][j] > maxvalue {
			maxvalue = states[n-1][j]
		}
	}
	return maxvalue
}
