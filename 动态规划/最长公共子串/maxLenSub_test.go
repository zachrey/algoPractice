package al42maxlen

import "testing"

func Test_swDP(t *testing.T) {
	a := []string{"m", "i", "t", "c", "m", "u"}
	b := []string{"m", "t", "a", "c", "n", "u"}

	n, m := 6, 6
	distance := getMaxLen(a, b, n, m)
	t.Log(distance)
}
