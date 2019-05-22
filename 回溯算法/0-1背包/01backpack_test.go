package al3901backpack

import "testing"

func Test_01backpack(t *testing.T) {
	n, w := 10, 100
	items := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 30}

	f(0, 0, items, n, w)

	t.Log(maxW)

	f2(0, 0, 0)
	t.Log(maxV)
}
