package sorting

func MergeSort(A []int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
}

func merge(A []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1)
	R := make([]int, n2)

	for i := range L {
		L[i] = A[p+i]
	}

	for j := range R {
		R[j] = A[p+1+j]
	}

	i := 0
	j := 0
	k := p

	for i < n1 || j < n2 {
		if L[i] <= R[j] {
			A[k] = L[i]
			k++
			i++
		} else {
			A[k] = R[j]
			k++
			j++
		}
	}
}
