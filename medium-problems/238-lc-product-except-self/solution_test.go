package lcproductexceptself

import "testing"

func TestReverseVowels(t *testing.T) {

	t.Run("[1,2,3,4]", func(t *testing.T) {
		got := productExceptSelf([]int{1, 2, 3, 4})
		want := []int{24, 12, 8, 6}

		for i, v := range want {
			if v != got[i] {
				t.Errorf("got %q, wanted %q", got, want)
			}
		}
	})

	t.Run("[-1,1,0,-3,3]", func(t *testing.T) {
		got := productExceptSelf([]int{-1, 1, 0, -3, 3})
		want := []int{0, 0, 9, 0, 0}

		for i, v := range want {
			if v != got[i] {
				t.Errorf("got %q, wanted %q", got, want)
			}
		}
	})
}
