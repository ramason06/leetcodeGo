package medium

func jump(pos int, nums []int, visited []bool) bool {
	if pos == len(nums)-1 {
		return true
	}

	if visited[pos] != false {
		return visited[pos]
	}

	maxJump := min(pos+nums[pos], len(nums)-1)
	for nextPos := pos + 1; nextPos <= maxJump; nextPos++ {
		if jump(nextPos, nums, visited) {
			visited[pos] = true
			return true
		}
	}

	visited[pos] = false
	return false
}

func jumpGreedy(nums []int) bool {
	maxReach := 0
	for i := 0; i <= maxReach; i++ {
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func canJump(nums []int) bool {
	visited := make([]bool, len(nums))
	return jump(0, nums, visited)
}
