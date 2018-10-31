package sorting

func QuickSort(values []int, left int, right int) {
	if right <= left {
		return
	}

	pivot := values[left]
	i := left + 1

	for j := left; j <= right; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[left], values[i-1] = values[i-1], pivot

	QuickSort(values, left, i-2)
	QuickSort(values, i, right)
}
