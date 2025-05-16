package merge2strings

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var str string
	for i, j := 0, 0; i < len(word1) || j < len(word2); {
		// str = str + string(word1[i]) + string(word2[j])
		if i < len(word1) {
			str = str + string(word1[i])
			i++
		}
		if j < len(word2) {
			str = str + string(word2[j])
			j++
		}

	}
	return str
}

func runMergeAlternatively() {
	result := mergeAlternately("abc", "pqr")
	// result := mergeAlternately("ab", "pqrs")
	fmt.Println("result=", result)
}
