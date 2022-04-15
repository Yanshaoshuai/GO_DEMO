package main

//递归
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
//}
//
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//队列 循环
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	l := 0
	var node *TreeNode
	for len(queue) != 0 {
		l++
		size := len(queue)
		for i := 0; i < size; i++ {
			node = queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return l
}
