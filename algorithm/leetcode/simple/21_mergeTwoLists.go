package main

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root = new(ListNode)
	var cursor = new(ListNode)
	cursor = root
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cursor.Next = list2
			cursor = cursor.Next
			list2 = list2.Next
		} else {
			cursor.Next = list1
			cursor = cursor.Next
			list1 = list1.Next
		}
	}
	if list1 != nil {
		cursor.Next = list1
	}
	if list2 != nil {
		cursor.Next = list2
	}
	return root.Next
}
