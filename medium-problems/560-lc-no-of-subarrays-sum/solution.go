package lcnoofsubarrayssum

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0
	currentSum := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if prefixSum[currentSum-k] > 0 {
			count += prefixSum[currentSum-k]
		}
		prefixSum[currentSum]++
	}
	return count
}

func main() {
	// input := []int{1, 1, 1}
	// k := 2
	// input := []int{1, 2, 3}
	// k := 3

	// input := []int{10, 2, -2, -20, 10} // output = 3
	// k := -10

	input := []int{1, -1, 0}
	k := 0

	fmt.Println("output ==> ", subarraySum(input, k))
}
