package binarytree

func (t *TreeNode) search(target int) *TreeNode {
	tt := t
	for tt != nil {
		if tt.Val == target {
			return tt
		} else if tt.Val > target {
			tt = tt.Left
		} else {
			tt = tt.Right
		}
	}

	return nil
}
