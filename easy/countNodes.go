package easy

import "leetcodeGo/common"

func count(pos *common.TreeNode) int {
	if pos.Left != nil {
		if pos.Right != nil {
			return count(pos.Left) + count(pos.Right) + 1
		} else {
			return count(pos.Left) + 1
		}
	}
	return 1
}

func countNodes(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root)
}
