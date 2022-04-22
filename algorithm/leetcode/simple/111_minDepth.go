package main

import "math"

//最小高度<==>第一个左右子树都为null的节点的高度

//递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return help(root)
}
func help(root *TreeNode) int {
	if root == nil { //剪枝
		return math.MaxInt
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return 1 + min(help(root.Left), help(root.Right))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//队列
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var queue []*TreeNode
//	var h int
//	queue = append(queue, root)
//	l := len(queue)
//	for l != 0 {
//		h++
//		for i := 0; i < l; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left == nil && node.Right == nil {
//				return h
//			}
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//		l = len(queue)
//	}
//	return h
//}
