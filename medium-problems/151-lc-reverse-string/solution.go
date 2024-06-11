package lcreversestring

import "fmt"

func isAlphabet(ch byte) bool {
	if ch >= 65 && ch <= 90 {
		return true
	}
	if ch >= 97 && ch <= 122 {
		return true
	}
	if ch >= 48 && ch <= 57 {
		return true
	}
	return false
}

func reverseStr(s string) string {
	var result []byte
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return string(result)
}

func removeTrailingSpace(s string) string {
	// var result []byte
	lastAlphabetIndex := 0
	for i := 0; i < len(s); i++ {
		if isAlphabet(s[i]) {
			lastAlphabetIndex = i
		}
	}
	if lastAlphabetIndex != (len(s) - 1) {
		return s[:lastAlphabetIndex+1]
	}
	return s
}

/*
*
f--> first pointer
s--> second pointer
"  hello world  "

	^^
	fs
*/
func reverseWords(s string) string {
	var output string

	tempStr := ""
	first, second := len(s)-2, len(s)-1
	for second >= 0 {
		if isAlphabet(s[second]) {
			tempStr += string(s[second])
			if first == -1 || !isAlphabet(s[first]) {
				// time to create word from temp string
				tempSpace := " "
				if second == 0 {
					tempSpace = ""
				}
				output += reverseStr(tempStr) + tempSpace
				tempStr = ""
			}
		}
		first--
		second--
	}
	return removeTrailingSpace(output)
}

func main() {
	// s := "the sky is blue"
	// s := "  hello world  "
	s := "EPY2giL"

	fmt.Println("received ===> \n", reverseWords(s))
}
