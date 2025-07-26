package easy

func moveZeroes(nums []int) {
	var cnt = 0
	for index, num := range nums {
		if num != 0 {
			nums[cnt], nums[index] = nums[index], nums[cnt]
			cnt++
		}
	}
}
