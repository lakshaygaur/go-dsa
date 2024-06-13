package lcproductexceptself

import "fmt"

/*
*
input : [1, 2, 3, 4]
left product  : 1 | 1 | 1*2 | 1*2*3
right product : 2*3*4 | 3*4 | 4 | 1
final product : ( left product * right product ) for each element
*/
func productExceptSelf(nums []int) []int {
	var output []int
	var leftProducts []int
	left, right := 1, len(nums)-2

	temp := 1
	leftProducts = append(leftProducts, temp)
	for ; left < len(nums); left++ {
		temp = temp * nums[left-1]
		leftProducts = append(leftProducts, temp)
	}

	temp = 1
	rightProducts := make([]int, len(nums))
	rightProducts[right+1] = temp
	for ; right >= 0; right-- {
		temp = temp * nums[right+1]
		rightProducts[right] = temp
	}
	// build final output
	for i := 0; i < len(nums); i++ {
		output = append(output, leftProducts[i]*rightProducts[i])
	}
	return output
}

func main() {
	input := [4]int{1, 2, 3, 4}

	fmt.Println("Result ==> ", productExceptSelf(input[:]))
}
