package easy

func singleNumber(nums []int) int {
	m := make(map[int]bool)

	for _, val := range nums {
		if m[val] == true {
			delete(m, val)
		} else {
			m[val] = true
		}
	}

	for k, _ := range m {
		return k
	}
	return 0
}
