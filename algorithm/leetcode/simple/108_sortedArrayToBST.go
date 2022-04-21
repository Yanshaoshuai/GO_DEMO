package main

//分治法 todo 循环
//root[0,n]={root[0,mid-1],num(mid),root[mid-1,n]}
//root[left,right]={root[left,mid-1],num(mid),root[mid-1,right]}
func sortedArrayToBST(nums []int) *TreeNode {
	return recursion(nums, 0, len(nums)-1)
}

func recursion(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = recursion(nums, left, mid-1)
	root.Right = recursion(nums, mid+1, right)
	return root
}
