package easy

import (
	"leetcodeGo/common"
	"math"
)

func getMinimumDifference(root *common.TreeNode) int {
	prev := -1
	minDiff := math.MaxInt32

	var inorder func(node *common.TreeNode)
	inorder = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != -1 {
			diff := node.Val - prev
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}
