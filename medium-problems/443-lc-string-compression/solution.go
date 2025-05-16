// package lcstringcompression
package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	// newLength := 0
	var tempInput string //[]byte
	charMap := make(map[string]int)
	// newCharToAdd := chars[0]
	// flag := true
	for i := 0; i < len(chars); i++ {
		// if flag {
		// 	tempInput = append(tempInput, newCharToAdd) // add newCharacter only if the flag allows
		// 	flag = false
		// }
		// start counting length of the character
		charMap[string(chars[i])]++
	}
	for key, value := range charMap {
		tempInput += key
		fmt.Println("value", value)
		tempInput += strconv.Itoa(value)
	}
	chars = []byte(tempInput) // should point to new tempInput variable
	fmt.Println("tempInput", tempInput, len(tempInput))
	fmt.Println("chars new ", string(chars))
	// chars[0] = 'b' // should point to new tempInput variable
	return len(tempInput)
}

func main() {
	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}

	fmt.Println("Output ===> ", compress(input))
	fmt.Println("Updated input ===> ", string(input))
}
