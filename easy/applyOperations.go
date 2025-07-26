package easy

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}

func applyOperationsAI(nums []int) []int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		// 연산 적용
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}

		// 동시에 0이 아닌 값 앞으로 복사
		if nums[i] != 0 {
			nums[pos] = nums[i]
			if pos != i {
				nums[i] = 0
			}
			pos++
		}
	}

	return nums
}
