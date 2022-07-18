package linkedlist

// partition LeetCode 86 分离链表
func partition(l *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	p1 := dummy1
	p2 := dummy2
	for l != nil {
		if l.Val < x {
			p1.Next = l
			p1 = p1.Next
		} else {
			p2.Next = l
			p2 = p2.Next
		}
		l = l.Next
	}
	p2.Next = nil
	p1.Next = dummy2.Next

	return dummy1.Next
}
