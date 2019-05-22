package al40triangle

import "testing"

var (
	items = [][]int{{5}, {7, 8}, {2, 3, 4}, {4, 9, 6, 1}, {2, 7, 9, 4, 5}}
	n     = 5
)

func Test_triangle(t *testing.T) {
	minV := getMinValue(items, n)
	t.Log(minV)
}
