package lcinreasingtripletsequence

import "testing"

func TestIncreasingTriplet(t *testing.T) {

	t.Run("[1,2,3,4,5]", func(t *testing.T) {
		got := increasingTriplet([]int{1, 2, 3, 4})
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("5, 4, 3, 2, 1", func(t *testing.T) {
		got := increasingTriplet([]int{5, 4, 3, 2, 1})
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("[2, 1, 5, 0, 4, 6]", func(t *testing.T) {
		got := increasingTriplet([]int{2, 1, 5, 0, 4, 6})
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("[20, 100, 10, 12, 5, 13]", func(t *testing.T) {
		got := increasingTriplet([]int{20, 100, 10, 12, 5, 13})
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("[1, 2, 2, 1]", func(t *testing.T) {
		got := increasingTriplet([]int{1, 2, 2, 1})
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("[1, 2, 1, 3]", func(t *testing.T) {
		got := increasingTriplet([]int{1, 2, 1, 3})
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("[1, 1, -2, 6]", func(t *testing.T) {
		got := increasingTriplet([]int{1, 1, -2, 6})
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("[1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]", func(t *testing.T) {
		got := increasingTriplet([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})
}
