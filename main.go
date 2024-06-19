package main

import "fmt"

/**
large array of int
given a large array of integers and an integer k, find the number of subarrays whose sum equals k.
Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2
--------------------------------------------------------------
Concurrent Prime Number Finder Write a Go program that finds all prime numbers up to a given number N using concurrency. The program should divide the range of numbers among multiple Go routines to check for primality and then collect the results through channels.
*/

// func sum(arr []int) int {
// 	result := 0
// 	for _, v := range arr {
// 		result += v
// 	}
// 	return result
// }

func subArray(input []int, k int) int {
	result := 0
	count := 0
	// var subarr []int
	for i := 0; i < len(input); i++ {
		count += input[i]
		// subarr
		if count == k {
			// reset count
			count = input[i]
			// increment result
			result++
		}
		if input[i] == k {
			result++
		}
	}
	return result
}

func main() {
	// input := []int{1, 1, 1}
	// k := 2
	// input := []int{1, 2, 3}
	// k := 3

	input := []int{10, 2, -2, -20, 10} // output = 3
	k := -10

	fmt.Println("output ==> ", subArray(input, k))
}
