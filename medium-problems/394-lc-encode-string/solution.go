package lcencodestring

import (
	"strconv"
)

func isNumber(ch byte) bool {
	nums := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < len(nums); i++ {
		if string(ch) == nums[i] {
			return true
		}
	}
	return false
}

func createSubstr(count int, s string) string {
	tempStr := ""
	for i := 0; i < count; i++ {
		tempStr += s
	}
	return tempStr
}

func decodeString(s string) string {
	output := ""
	// stack based approach
	/**
	| 3 | // store in num
	| [ | // start putting in a stack; wait for ] to close substring
	| a | // a store in substring
	| 2 | // 2 store in num
	| [ | // start putting in a stack; wait for ] to close substring
	| c | // c
	| ] | // close substring
	| ] | // close substring
	*/
	var nums []int // number stack
	var strStack []string
	k := 0

	for i := 0; i < len(s); i++ {
		if isNumber(s[i]) {
			// store in num stack
			temp, _ := strconv.Atoi(string(s[i]))
			k = (k * 10) + temp
			continue
		}

		if string(s[i]) == "[" {
			nums = append(nums, k)
			k = 0
			// store characters in string stack
			strStack = append(strStack, string(s[i]))
			continue
		}

		if string(s[i]) != "]" {
			// store characters in string stack
			strStack = append(strStack, string(s[i]))
			continue
		}

		if string(s[i]) == "]" {
			// start popping
			var substr []string
			j := len(strStack) - 1
			for ; string(strStack[j]) != "["; j-- {
				// substr += strStack[j]
				substr = append(substr, strStack[j])
			}
			// j--
			// pop the elements up until j in strStack
			strStack = strStack[:j]
			// reverse substr to form correct string
			tempStr := ""
			for m := len(substr) - 1; m >= 0; m-- {
				tempStr += string(substr[m])
			}
			// build string from substr and
			// push back in strStack
			strStack = append(strStack, createSubstr(nums[len(nums)-1], tempStr))
			// pop from num stack
			nums = nums[:len(nums)-1]
		}
	}
	for i := 0; i < len(strStack); i++ {
		output += strStack[i]
	}
	return output
}
