package easy

import "leetcodeGo/common"

func sortedArrayToBST(nums []int) *common.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return makeTree(nums, 0, len(nums)-1)
}
