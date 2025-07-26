package easy

import "leetcodeGo/common"

func avgOfLevels(root *common.TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var ans []float64
	var queue []*common.TreeNode

	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		avg := float64(levelSum) / float64(levelSize)
		ans = append(ans, avg)
	}

	return ans
}
