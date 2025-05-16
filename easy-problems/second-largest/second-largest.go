package secondlargest

import (
	"fmt"
)

// Find the 2nd largest number in array
// Arr = [2,3,45,1,10,100,71,55,45,0,-1, 3, 4]

func findSecond(arr []int) int {
	largest := 0
	second := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < largest && arr[i] > second {
			second = arr[i]
		}
	}
	fmt.Println("largest", largest)
	return second
}

func runFindSecond() {
	arr := []int{2, 3, 45, 1, 10, 100, 71, 55, 45, 0, -1, 3, 4}
	result := findSecond(arr)
	fmt.Println("second largest ", result)

}
