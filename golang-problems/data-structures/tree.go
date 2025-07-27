package ds

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func same(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return same(p.Left, q.Left) && same(p.Right, q.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return same(p, q)
}
