package binarytree

// 根据根节点左右划分子树，递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 左-右-根
func postOrder(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// check
		node := stack[len(stack)-1]
		// right-root
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// poped
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

func fansIn(chs ...chan interface{}) chan interface{} {
	c := make(chan interface{})
	for _, ch := range chs {
		ch := ch
		go func(in chan interface{}) {
			for {
				c <- <-ch
			}
		}(ch)
	}
	return c
}
