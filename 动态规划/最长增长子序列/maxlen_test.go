package al42maxlen

import "testing"

func Test_maxlen(t *testing.T) {
	a := []int{2, 9, 3, 6, 5, 1, 7}
	n := len(a)

	maxlen := getMaxLen(a, n)

	t.Log(maxlen)
}
