package lcnoofsubarrayssum

import "testing"

func TestSubarraySum(t *testing.T) {

	t.Run("[1, 1, 1]", func(t *testing.T) {
		input := []int{1, 1, 1}
		k := 2
		got := subarraySum(input, k)
		want := 2

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("[1, 2, 3]", func(t *testing.T) {
		input := []int{1, 2, 3}
		k := 3
		got := subarraySum(input, k)
		want := 2

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("[10, 2, -2, -20, 10]", func(t *testing.T) {
		input := []int{10, 2, -2, -20, 10} // output = 3
		k := -10
		got := subarraySum(input, k)
		want := 3

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

}
