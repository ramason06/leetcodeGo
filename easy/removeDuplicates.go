package easy

func removeDuplicates(nums []int) int {
	i := 0
	c := 0

	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
		} else {
			j := 0
			for i-j > c {
				nums[i-j] = nums[i-j+1]
				j++

			}
			c++
		}
		i++
	}

	return c + 1
}
