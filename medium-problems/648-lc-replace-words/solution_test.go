package lcreplacewords

import (
	"testing"
)

func TestReplaceWords(t *testing.T) {

	t.Run("Find root of sentence: the cattle was rattled by the battery", func(t *testing.T) {
		in := "the cattle was rattled by the battery"
		dictionary := []string{"cat", "bat", "rat"}
		got := replaceWords(dictionary, in)
		want := "the cat was rat by the bat"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
