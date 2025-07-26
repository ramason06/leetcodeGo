package easy

func majorityElement(nums []int) int {
	var am map[int]int
	am = make(map[int]int)

	for _, n := range nums {
		am[n]++
		if am[n] > len(nums)/2 {
			return n
		}
	}

	return -1
}
