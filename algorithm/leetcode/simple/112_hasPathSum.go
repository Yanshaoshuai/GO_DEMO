package main

//func hasPathSum(root *TreeNode, targetSum int) bool {
//	var queue []*TreeNode
//	var sum int
//	queue = append(queue, root)
//	l := 1
//	for l != 0 {
//		for i := 0; i < l; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//	}
//	return false
//}
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}
