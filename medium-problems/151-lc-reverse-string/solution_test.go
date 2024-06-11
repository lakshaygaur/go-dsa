package lcreversestring

import "testing"

func TestReverseWords(t *testing.T) {

	t.Run("the sky is blue", func(t *testing.T) {
		got := reverseWords("the sky is blue")
		want := "blue is sky the"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("  hello world  ", func(t *testing.T) {
		got := reverseWords("  hello world  ")
		want := "world hello"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("EPY2giL", func(t *testing.T) {
		got := reverseWords("EPY2giL")
		want := "EPY2giL"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
