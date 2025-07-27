package lcsametree

import "testing"

func TestSameTree(t *testing.T) {
	t.Run("p = [1,2,3], q = [1,2,3]", func(t *testing.T) {
		want := true
		p := &TreeNode{}
		p.Val = 1 // root
		p.Left = &TreeNode{Val: 2}
		p.Right = &TreeNode{Val: 3}
		q := &TreeNode{}
		q.Val = 1 // root
		q.Left = &TreeNode{Val: 2}
		q.Right = &TreeNode{Val: 3}

		got := isSameTree(p, q)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("p = [1,2], q = [1,null,2]", func(t *testing.T) {
		want := false
		p := &TreeNode{}
		p.Val = 1 // root
		p.Left = &TreeNode{Val: 2}

		q := &TreeNode{}
		q.Val = 1 // root
		q.Right = &TreeNode{Val: 2}

		got := isSameTree(p, q)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
