package binarytree

func maxDepth(root *TreeNode) int {
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
