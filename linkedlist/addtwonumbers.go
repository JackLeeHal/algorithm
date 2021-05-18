package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r *ListNode
	dummy := r
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1+n2+carry
		sum, carry = sum%10, sum/10
		dummy = &ListNode{Val: sum}
		dummy = dummy.Next
	}
	if carry > 0 {
		dummy.Next = &ListNode{Val: carry}
	}

	return r
}
