package lc

func InorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}
