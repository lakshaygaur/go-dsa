// Given two words (beginWord and endWord) and a word list, find the length of the shortest transformation sequence from beginWord to endWord, such that:
//  Comment

// only one letter can be changed at a time
// each transformed word must exist in the word list
// Example
// Input
// beginWord: "hit"
// endWord: cog
// wordList: ["hot", "dog", "dot", "lot", "log", "cog"]

// Output
// 5

// Transformation sequence
// "hit" --> "hot" --> "dot" --> "dog" --> "cog"

// Note
// return 0 if there is no such transformation sequence
// all words have the same length
// all words contain only lowercase alphabetic characters
// you may assume no duplicates
// you may assume beginWord and endWord are not the same word
package main

/**
hit, pat
iteration -1 ---> hit
{
	"hot": 1,
	"hat": 1,
	"dog": 0,
	"dot", "lot", "log", "cog", "cat"
}

iteration  - 2 --> [hat, hot]

*/

func diffInStr(s1 string, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	if diff > 1 {
		return false
	}
	return true
}

func transformSeq(input []string, beginWord string, endWord string) int {
	output := 0
	hashMap := make(map[string]int)
	for i := 0; i < len(input); i++ { // initialise
		hashMap[input[i]] = 0
	}
	iterationCount := 0
	word := beginWord
	for iterationCount < len(input) {
		for i := 0; i < len(input); i++ {
			// increment matched words
			if diffInStr(input[i], word) {
				hashMap[input[i]] = iterationCount + 1
			}
		}
		// re-analyse or select the word to match against

	}
	return output
}

// func main() {
// 	input := []string{"hot", "dog", "dot", "lot", "log", "cog"}
// 	beginWord := "hit"
// 	endWord := "cog"
// 	transformSeq(input, beginWord, endWord)

// }
