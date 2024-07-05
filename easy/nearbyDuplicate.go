package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if j, ex := m[v]; ex {
			if i-j <= k {
				return true
			}
		}
		m[v] = i
	}

	return false
}
