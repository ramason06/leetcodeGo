package easy

import "leetcodeGo/common"

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
