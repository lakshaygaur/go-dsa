package reverse_vowel

import "testing"

func TestReverseVowels(t *testing.T) {

	got := reverseVowels("leetcode")
	want := "leotcede"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
