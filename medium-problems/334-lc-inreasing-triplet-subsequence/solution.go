package lcinreasingtripletsequence

import (
	"math"
)

func increasingTriplet(nums []int) bool {

	first := int(math.Pow(2, 32))  //2 ^ 32
	second := int(math.Pow(2, 32)) //2 ^ 32

	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second { // && second > first
			second = nums[i]
		} else {
			return true
		}

	}

	return false
}
