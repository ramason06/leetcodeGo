package easy

import (
	"fmt"
	"leetcodeGo/common"
)

func preorderTraversal(root *common.TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}
