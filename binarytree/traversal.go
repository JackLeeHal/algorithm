package binarytree

import "fmt"

// 前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
// 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
// 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func preOrder2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			fmt.Println(root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
}

func inorder(root *TreeNode) {
	if root != nil {
		return
	}

	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(val.Val)
		root = val.Right
	}
}
