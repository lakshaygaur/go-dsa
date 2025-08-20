package lcissubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	t.Run("s = abc, t = axbyc", func(t *testing.T) {
		want := true
		got := isSubsequence("abc", "axbyc")

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
