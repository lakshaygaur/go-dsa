package lcsearchinbst

import "testing"

func TestSearchInBST(t *testing.T) {

	t.Run("the sky is blue", func(t *testing.T) {
		root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 7}}
		got := searchBST(root, 2)
		want := 2
		if got == nil || got.Val != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
