package main

//root is balanced <==>|H(root.Left)-H(root.Right)|<=1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := high(root.Left) - high(root.Right)
	return (-1 <= diff && 1 >= diff) && isBalanced(root.Left) && isBalanced(root.Right)
}
func high(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(high(root.Left), high(root.Right))
}
