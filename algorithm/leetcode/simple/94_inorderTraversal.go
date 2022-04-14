package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
//func inorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	return recursion(result, root)
//}
//
//func recursion(result []int, root *TreeNode) []int {
//	if root == nil {
//		return result
//	}
//	result = recursion(result, root.Left)
//	result = append(result, root.Val)
//	result = recursion(result, root.Right)
//	return result
//}
//栈
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		if root != nil {
			//把每一个左节点作为根节点 直到没有左节点
			stack = append(stack, root)
			root = root.Left
		} else {
			//从最左边开始 访问root节点 -->访问右节点
			//访问root节点
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, bottom.Val)
			//访问右节点
			root = bottom.Right
		}
	}
	return result
}
func main() {
	root := &TreeNode{Left: nil, Right: &TreeNode{Left: &TreeNode{Val: 3}, Val: 2}, Val: 1}
	result := inorderTraversal(root)
	fmt.Printf("%v", result)
}
