package main

//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	stack1 := make([]*TreeNode, 0)
//	stack2 := make([]*TreeNode, 0)
//	for (p != nil && q != nil) || (len(stack1) != 0 && len(stack2) != 0) {
//		if p != nil && q == nil || p == nil && q != nil {
//			return false
//		}
//		if p != nil && q != nil {
//			stack1 = append(stack1, p)
//			p = p.Left
//			stack2 = append(stack2, q)
//			q = q.Left
//		} else {
//			p = stack1[len(stack1)-1]
//			stack1 = stack1[:len(stack1)-1]
//			q = stack2[len(stack2)-1]
//			stack2 = stack2[:len(stack2)-1]
//			if p.Val != q.Val {
//				return false
//			}
//			p = p.Right
//			q = q.Right
//		}
//	}
//	return p == q
//}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	if p != nil {
		return isSameTree(p.Left, q.Left) && (p.Val == q.Val) && isSameTree(p.Right, q.Right)
	}
	return true
}
