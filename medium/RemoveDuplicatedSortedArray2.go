package medium

func removeDuplicates(nums []int) int {
	checker := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		if checker[nums[i]] < 2 {
			checker[nums[i]]++
			ans++
		} else if i+1 < len(nums) {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return ans
}

func removeDuplicatesAI(nums []int) int {
	checker := make(map[int]int)
	write := 0
	for read := 0; read < len(nums); read++ {
		val := nums[read]
		if checker[val] < 2 {
			checker[val]++
			nums[write] = val
			write++
		}
	}
	return write
}
