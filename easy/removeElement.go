package easy

func removeElement(nums []int, val int) int {
	i, k := 0, 0

	for i < len(nums) {
		if nums[i] != val {
			i++
			k++
		} else if i < len(nums)-1 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return k
}

func removeElement2(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
