package easy

import "leetcodeGo/common"

func makeTree(nums []int, left int, right int) *common.TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	node := &common.TreeNode{Val: nums[mid]}
	node.Left = makeTree(nums, left, mid+1)
	node.Right = makeTree(nums, mid, right)

	return node
}
