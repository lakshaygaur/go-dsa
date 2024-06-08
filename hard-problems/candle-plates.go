package main

import "fmt"

func countPlates(s string, query []int) int {
	// s := "*|***||*"
	count := 0
	result := 0
	flag := false
	startIndex := query[0]
	endIndex := query[1]
	// var result []int
	for i := startIndex; i <= endIndex; i++ {
		// fmt.Printf("%c", s[i])
		// test := s[i] == '|'
		// fmt.Println(test)
		if s[i] == '|' {
			if flag == true {
				if count > result {
					result = count
				}
				flag = false
			} else {
				// fmt.Printf("%c", s[i])
				// fmt.Println("increasing start")
				flag = true
			}
		}
		if flag && s[i] == '*' {
			// increase count
			count++
		}
	}
	return count
}

func main() {
	// s := "*|***||*"
	s := "**|**|***|"
	var query []int
	query = append(query, 2)
	query = append(query, 5)
	result := countPlates(s, query)
	fmt.Println(result)
}
