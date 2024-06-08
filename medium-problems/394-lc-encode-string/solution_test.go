package lcencodestring

import "testing"

func TestReverseVowels(t *testing.T) {

	t.Run("3[a]2[bc]", func(t *testing.T) {
		got := decodeString("3[a]2[bc]")
		want := "aaabcbc"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("3[a2[c]]", func(t *testing.T) {
		got := decodeString("3[a2[c]]")
		want := "accaccacc"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("3[z]2[2[y]pq4[2[jk]e1[f]]]ef", func(t *testing.T) {
		got := decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
		want := "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
