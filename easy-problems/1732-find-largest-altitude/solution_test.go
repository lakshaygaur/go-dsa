package findlargestalt

import "testing"

func TestLargestAltitude(t *testing.T) {

	t.Run("[-5,1,5,0,-7]", func(t *testing.T) {
		got := largestAltitude([]int{-5, 1, 5, 0, -7})
		want := 1

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
