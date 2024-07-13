package easy

import "leetcodeGo/common"

func sumTree(s *common.TreeNode, targetSum int, posSum int) bool {
	if s == nil {
		return false
	}
	if s.Left == nil && s.Right == nil {
		if posSum+s.Val == targetSum {
			return true
		}
	}
	return sumTree(s.Left, targetSum, posSum+s.Val) || sumTree(s.Right, targetSum, posSum+s.Val)
}

func hasPathSum(root *common.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return sumTree(root, targetSum, 0)
}
