package linkedlist

// 判断环形链表
func hasCycle(head *ListNode) bool {
	// 思路：快慢指针 快慢指针相同则有环，证明：如果有环每走一步快慢指针距离会减1
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		// 比较指针是否相等（不要使用val比较！）
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
