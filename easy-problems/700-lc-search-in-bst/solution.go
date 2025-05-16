package lcsearchinbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	tempRoot := root
	for tempRoot != nil {
		if tempRoot.Val == val {
			return tempRoot
		}
		if tempRoot.Val < val { // go right
			tempRoot = tempRoot.Right
			continue
		}
		if tempRoot.Val > val { // go left
			tempRoot = tempRoot.Left
		}
	}
	return nil
}
