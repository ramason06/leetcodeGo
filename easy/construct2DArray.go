package easy

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return make([][]int, 0)
	}

	idx := 0
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr[i][j] = original[idx]
			idx++
		}
	}

	return arr
}
