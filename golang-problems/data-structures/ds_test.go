package ds

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	head := Constructor()
	head.Print()
	t.Run("Add at head", func(t *testing.T) {
		head.AddAtHead(1)
		head.Print()
	})

	t.Run("Add at tail", func(t *testing.T) {
		head.AddAtTail(3)
		head.Print()
	})

	t.Run("Add at index", func(t *testing.T) {
		head.AddAtIndex(1, 2)
		head.Print()
	})

	t.Run("Get at index", func(t *testing.T) {
		got := head.Get(1)
		want := 2

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("Delete at index", func(t *testing.T) {
		head.DeleteAtIndex(1)
		head.Print()
	})

	t.Run("Get at index", func(t *testing.T) {
		got := head.Get(1)
		want := 3

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestTrie(t *testing.T) {

	root := &TrieNode{children: make(map[rune]*TrieNode, 26)}
	t.Run("Insert words", func(t *testing.T) {
		root.Insert("bat")
		root.Insert("cat")
		root.Insert("rat")
		root.Print()
		fmt.Println("Inserting overlapping word")
		root.Insert("rap")
		root.Print()
	})

	t.Run("Find words: cat", func(t *testing.T) {
		got := root.Find("cat")
		want := true

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Find words: cab", func(t *testing.T) {
		got := root.Find("cab")
		want := false

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Find words: rap", func(t *testing.T) {
		got := root.Find("rap")
		want := true

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Find root: cattle", func(t *testing.T) {
		got := root.FindRoot("cattle")
		want := "cat"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

}

func TestReplaceWords(t *testing.T) {

	t.Run("Find root of sentence: the cattle was rattled by the battery", func(t *testing.T) {
		in := "the cattle was rattled by the battery"
		dictionary := []string{"cat", "bat", "rat"}
		got := replaceWords(dictionary, in)
		want := "the cat was rat by the bat"
		fmt.Println("")
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
