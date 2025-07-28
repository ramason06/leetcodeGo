package medium

func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	count := 0

	var dfs func(i int, or int)
	dfs = func(i int, or int) {
		if i == len(nums) {
			if or == maxOr {
				count++
			} else if or > maxOr {
				maxOr = or
				count = 1
			}
			return
		}

		dfs(i+1, or|nums[i])
		dfs(i+1, or)
	}
	dfs(0, 0)
	return count
}
