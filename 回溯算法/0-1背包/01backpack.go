package al3901backpack

import (
	"math"
)

/**
i 当前遍历的物品下标
cw 当前背包里面装了多少
items 每一个物品的重量
n 总共多少个物品
w 背包最多能装多少
*/

var maxW = math.MinInt16

func f(i, cw int, items []int, n, w int) {
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	f(i+1, cw, items, n, w) // 用来遍历 每个物品是否装的情况, 递归变成了到序的求解
	if cw+items[i] <= w {   // 如果装下的重量比背包的容量大，就不继续装了
		f(i+1, cw+items[i], items, n, w)
	}
}

var (
	items  = []int{2, 2, 4, 6, 3} // 每个物品的重量
	values = []int{3, 4, 8, 9, 6} // 每个物品的价值
	n, w   = 5, 9                 // 总共5个物品，背包容量为9
	maxV   = math.MinInt16
)

func f2(i, cw, cv int) {
	if cw == w || i == n {
		if cv > maxV {
			maxV = cv
		}
		return
	}
	f2(i+1, cw, cv)
	if cw+items[i] <= w {
		f2(i+1, cw+items[i], cv+values[i])
	}
}
