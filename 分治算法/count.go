package al38

var num = 0

// Count 计算逆序对个数
func Count(a []int, n int) int {
	mergeSortCounting(a, 0, n-1)
	return num
}

func mergeSortCounting(a []int, l, r int) {
	if l >= r {
		return
	}

	q := (l + r) / 2

	mergeSortCounting(a, l, q)
	mergeSortCounting(a, q+1, r)
	merge(a, l, q, r)
}

func merge(a []int, l, q, r int) {
	i, j, k := l, q+1, 0

	tmp := make([]int, r-l+1)
	for i <= q && j <= r {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			num += (q - i + 1) // 统计 左边 已经排好序的，比 右边 排好序的 大的数有多少， 这里为什么是 (q - i + 1) 而不是 1 可以仔细再想想
			tmp[k] = a[j]
			k++
			j++
		}
	}

	for i <= q {
		tmp[k] = a[i]
		k++
		i++
	}

	for j <= r {
		tmp[k] = a[j]
		k++
		j++
	}

	for i := 0; i <= r-l; i++ {
		a[l+i] = tmp[i]
	}
}
