// package lcinreasingtripletsequence
package main

import "fmt"

func increasingTriplet(nums []int) bool {

	i := 0
	j := i + 1
	k := j + 1

	for i < len(nums)-2 {
		fmt.Println("i=", i, " j=", j, " k=", k)
		if nums[i] < nums[j] && nums[j] < nums[k] {
			return true
		}
		if nums[i] > nums[j] {
			// increment i and j
			i++
			if i == j {
				j++
			}
			if j == k {
				k++
			}
			continue
		}
		if nums[j] > nums[k] {
			j++
			k++
			continue
		}
	}
	return false
}

func main() {
	// input := []int{1, 2, 3, 4, 5}
	// input := []int{5, 4, 3, 2, 1}
	// input := []int{2, 1, 5, 0, 4, 6}
	input := []int{20, 100, 10, 12, 5, 13}

	fmt.Println("Output ===> ", increasingTriplet(input))
}
