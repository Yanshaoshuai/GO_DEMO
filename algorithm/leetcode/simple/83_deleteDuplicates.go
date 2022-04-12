package main

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := head
	for dummy != nil && dummy.Next != nil {
		if dummy.Val == dummy.Next.Val {
			dummy.Next = dummy.Next.Next
			continue
		}
		dummy = dummy.Next
	}
	return head
}
