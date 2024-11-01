package lcmaxavgsubarray1

func findMaxAverage(nums []int, k int) float64 {
	left := 0
	right := k - 1
	var avg float64
	windowSum := 0
	// calulcate sum of initial window
	for i := left; i <= right; i++ {
		windowSum = windowSum + nums[i]
	}

	avg = float64(windowSum) / float64(k)
	sum := windowSum

	// slide the window and compare
	left++
	right++
	for right < len(nums) {

		sum = sum + nums[right] - nums[left-1]

		if sum > windowSum {
			avg = float64(sum) / float64(k)
			windowSum = sum
		}

		left++
		right++
	}
	return avg
}
