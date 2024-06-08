package reverse_vowel

import "fmt"

func isVowel(ch byte) bool {
	vowels := [10]byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for i := 0; i < len(vowels); i++ {
		if vowels[i] == ch {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	var output []byte
	// build replacable string
	for i := 0; i < len(s); i++ {
		output = append(output, s[i])
	}

	start := 0
	end := len(output) - 1
	for start < end {
		if isVowel(output[start]) && isVowel(output[end]) {
			//swap places
			temp := output[end]
			output[end] = output[start]
			output[start] = temp
			start++
			end--
		}
		// keep incrementing pointers if it is not a vowel
		if !isVowel(output[start]) {
			start++
		}
		if !isVowel(output[end]) {
			end--
		}
	}
	fmt.Println("output", string(output))
	return string(output)
}
