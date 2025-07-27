package gcdstrings

import "testing"

func TestSearchInBST(t *testing.T) {

	t.Run("the sky is blue", func(t *testing.T) {
		got := gcdOfStrings("ABCABC", "ABC")
		want := "ABC"
		if got == nil || got.Val != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
