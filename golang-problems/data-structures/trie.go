package ds

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

/*
	root -> b,	 c,	   r
			|	 | 	   |
			a    a      a
			|    |    |   |
			t    t    t   p
*/

// accepts a word to be inserted
func (p *TrieNode) Insert(word string) {
	cur := p
	for _, r := range word {
		if cur.children[r] == nil {
			cur.children[r] = &TrieNode{
				children: make(map[rune]*TrieNode, 26),
			}
		}
		cur = cur.children[r]
	}
	// go back
	cur.isWord = true
}

func (p *TrieNode) Print() {

	for k, _ := range p.children {
		fmt.Print(" ", string(k))
		p.children[k].Print()
	}
	fmt.Print("\n")
}

func (p *TrieNode) Find(word string) bool {
	cur := p
	for _, r := range word {
		_, ok := cur.children[r]
		if !ok {
			return false
		}
		cur = cur.children[r]
	}
	return cur.isWord
}

func (p *TrieNode) FindRoot(word string) string {
	cur := p
	str := ""
	for _, r := range word {
		_, ok := cur.children[r]
		// fmt.Println(", is word: ", cur.isWord)
		if cur.isWord {
			return str
		}
		if !ok {
			return word
		}
		// fmt.Println("adding:", r, " new str: ", str, ", is word: ", cur.isWord)
		str = str + "" + string(r)
		cur = cur.children[r]
	}
	return str
}

func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	var newSentence []string
	root := &TrieNode{
		children: make(map[rune]*TrieNode, 26),
	}
	for _, d := range dictionary {
		root.Insert(d)
	}

	for _, w := range words {
		newW := root.FindRoot(w)
		newSentence = append(newSentence, newW)
	}

	return strings.Join(newSentence, " ")

}
