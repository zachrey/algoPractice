package al38

import "testing"

func Test_Count(t *testing.T) {
	arr := []int{2, 4, 3, 1, 5, 6}
	count := Count(arr, len(arr))

	t.Log(count)
}
