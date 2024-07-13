package easy

import "leetcodeGo/common"

func isSymmetricFun(s1, s2 *common.TreeNode) bool {
	if s1 == nil && s2 == nil {
		return true
	}
	if s1 == nil || s2 == nil {
		return false
	}
	return (s1.Val == s2.Val) && isSymmetricFun(s1.Left, s2.Right) && isSymmetricFun(s1.Right, s2.Left)
}

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricFun(root.Left, root.Right)
}
