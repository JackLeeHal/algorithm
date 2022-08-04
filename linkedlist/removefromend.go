package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	p := findNode(dummy, n+1)
	p.Next = p.Next.Next

	return dummy.Next
}

func findNode(head *ListNode, n int) *ListNode {
	l1 := head
	for i := 0; i < n; i++ {
		l1 = l1.Next
	}
	l2 := head
	for l1 != nil {
		l1 = l1.Next
		l2 = l2.Next
	}

	return l2
}
