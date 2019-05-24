package al42lwst

import "testing"

func Test_lwst(t *testing.T) {
	a := []string{"m", "i", "t", "c", "m", "u"}
	b := []string{"m", "t", "a", "c", "n", "u"}

	n, m := 6, 6
	lwst(0, 0, n, m, 0, a, b)
	t.Log(edist)
}
