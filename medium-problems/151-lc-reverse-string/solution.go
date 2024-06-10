package main

import "fmt"

func isAlphabet(ch byte) bool {
	if ch >= 65 && ch <= 90 {
		return true
	}
	if ch >= 97 && ch <= 122 {
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

func reverseWords(s string) string {
	var output string
	// for i := 0; i < len(s); i++ {
	// 	output = append(output, s[i])
	// }

	tempStr := ""
	first, second := len(s)-2, len(s)-1
	for second >= 0 {
		fmt.Println("s[i]", string(s[second]), isAlphabet(s[second]))
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
	return output
}

func main() {
	// s := "the sky is blue"
	s := "  hello world  "
	// fmt.Println(" ascii ", 'z')

	fmt.Println("received ===> \n", reverseWords(s))
}
