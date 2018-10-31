package sorting

func CountingSort(array []int) {
	k := array[0]

	for _, val := range array {
		if k < val {
			k = val
		}
	}

	B := make([]int, len(array))
	copy(B, array)
	C := make([]int, k)

	for _, val := range array {
		C[val-1]++
	}

	for i := 1; i < k; i++ {
		C[i] += C[i-1]
	}

	for j := len(array) - 1; j > -1; j-- {
		array[C[B[j]-1]-1] = B[j]
		C[B[j]-1]--
	}
}
