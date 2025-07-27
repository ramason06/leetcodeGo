package easy

func countHillValley(nums []int) int {
	ans := 0
	lastVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] && lastVal < nums[i-1] {
			ans++
			lastVal = nums[i-1]
		} else if nums[i] > nums[i-1] && lastVal > nums[i-1] {
			ans++
			lastVal = nums[i-1]
		} else {

		}
	}
	return ans
}
