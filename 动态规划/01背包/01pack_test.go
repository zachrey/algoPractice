package al40pack

import "testing"

func Test_01backpack(t *testing.T) {
	n, w := 10, 100
	items := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 30}

	i := knapsack(items, n, w)

	i2 := knapsack2(items, n, w)

	t.Log(i)
	t.Log(i2)
	t.Log(knapsack3())
}
