package lc605

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	arr := []int{1, 0, 0, 0, 1}
	n := 2

	isCan := canPlaceFlowers(arr, n)

	t.Log(isCan)
}
