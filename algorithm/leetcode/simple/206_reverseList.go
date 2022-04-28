package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := head
	next := head.Next
	//此时dummy.Next=next并且不会被修改
	//所以手动置为null
	head.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = dummy
		dummy = next
		next = tmp
	}
	return dummy
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	root = reverseList(root)
	for root != nil {
		println(root.Val)
		root = root.Next
	}
}
