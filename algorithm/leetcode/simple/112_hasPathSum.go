package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	var stack []*TreeNode
	var sum int
	var lastVisit *TreeNode
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			sum += root.Val
			if root.Left == nil && root.Right == nil && sum == targetSum {
				return true
			}
			lastVisit = root
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			if node.Right != nil && lastVisit != node.Right {
				root = node.Right
			} else {
				if node.Left != nil {
					sum -= node.Left.Val
				}
				if node.Right != nil {
					sum -= node.Right.Val
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return false
}

//func hasPathSum(root *TreeNode, targetSum int) bool {
//	if root == nil {
//		return false
//	}
//	if root.Left == nil && root.Right == nil && targetSum == root.Val {
//		return true
//	}
//	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
//}
