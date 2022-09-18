package easy

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if v, existed := m[target-num]; existed {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}
