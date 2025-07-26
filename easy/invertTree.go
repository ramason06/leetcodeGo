package easy

import "leetcodeGo/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return root
	}
	right := root.Right
	left := root.Left

	tmp := right
	root.Right = left
	root.Left = tmp

	return invertTree(root.Right)
	return invertTree(root.Left)
}
