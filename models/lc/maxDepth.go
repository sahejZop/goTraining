package lc

import _ "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
