package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	rob := 0
	norob := 0
	for i := 0; i < len(nums); i++ {
		newRob := norob + nums[i]
		newNoRob := max(norob, rob)
		rob = newRob
		norob = newNoRob
		fmt.Println("rob=", rob, " norob=", norob)
	}
	return max(rob, norob)
}
func main() {
	// input := []int{2, 1, 1, 2}
	input := []int{1, 2, 3, 1}
	fmt.Println("result : ", rob(input))
}
