package lc605

func canPlaceFlowers(flowers []int, n int) bool {
	tmp := append([]int{0}, flowers...)
	tmp = append(tmp, 0)

	for i := 1; i < len(tmp)-1; i++ {
		if tmp[i-1] == 0 && tmp[i] == 0 && tmp[i+1] == 0 {
			n--
		}
	}
	if n <= 0 {
		return true
	}
	return false
}
