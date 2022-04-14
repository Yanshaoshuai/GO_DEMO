package main

//递归
//func isSymmetric(root *TreeNode) bool {
//	if root.Left != nil && root.Right == nil || root.Right != nil && root.Left == nil {
//		return false
//	}
//	return recursion(root.Left, root.Right)
//}
//
//func recursion(left *TreeNode, right *TreeNode) bool {
//	if left == nil && right != nil || left != nil && right == nil {
//		return false
//	}
//	if left != nil && right != nil {
//		return left.Val == right.Val && recursion(left.Left, right.Right) && recursion(left.Right, right.Left)
//	}
//	return true
//}

//栈 循环
func isSymmetric(root *TreeNode) bool {
	if root.Left != nil && root.Right == nil || root.Right != nil && root.Left == nil {
		return false
	}
	p := root.Left
	q := root.Right
	if p != nil && q != nil {
		stack1 := make([]*TreeNode, 0)
		stack2 := make([]*TreeNode, 0)
		for p != nil && q != nil || len(stack1) != 0 && len(stack2) != 0 {
			if p != nil && q == nil || p == nil && q != nil {
				return false
			}
			if p != nil && q != nil {
				stack1 = append(stack1, p)
				stack2 = append(stack2, q)
				p = p.Left
				q = q.Right
			} else {
				p = stack1[len(stack1)-1]
				q = stack2[len(stack2)-1]
				stack1 = stack1[:len(stack1)-1]
				stack2 = stack2[:len(stack2)-1]
				if p.Val != q.Val {
					return false
				}
				p = p.Right
				q = q.Left
			}
		}
	}
	return p == q
}
