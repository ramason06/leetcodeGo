package easy

func findIndex(nums []int, left int, right int, target int) int {
	if left > right {
		return left
	}

	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return findIndex(nums, left, mid-1, target)
	} else {
		return findIndex(nums, mid+1, right, target)
	}
}

func searchInsert(nums []int, target int) int {
	return findIndex(nums, 0, len(nums)-1, target)
}
