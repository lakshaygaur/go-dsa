package mergesort

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	mIndex := m - 1
	nIndex := n - 1

	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1[:n], nums2)
		return
	}

	for i := m + n - 1; i >= 0; i-- {
		fmt.Print("first iteration ", nIndex, " m ", mIndex)
		if nIndex >= 0 {
			if nums2[nIndex] > nums1[mIndex] {
				nums1[i] = nums2[nIndex]
				nIndex--
			}
		} else if mIndex >= 0 {
			nums1[i] = nums1[mIndex]
			mIndex--
		}
		fmt.Print(" a [i] = ", i, "\t")
		for j := 0; j < len(nums1); j++ {
			fmt.Print(nums1[j], ",")
		}
		fmt.Print("\t nIndex=", nIndex, ", mIndex=", mIndex)
		fmt.Print("\n")
	}
	// fmt.Print("Result = ")
	// for i := 0; i < len(nums1); i++ {
	// 	fmt.Print(nums1[i], ",")
	// }
}

func runMerge() {

	//case 1
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3

	//case 2
	// nums1 := []int{1}
	// m := 1
	// nums2 := []int{}
	// n := 0

	//case 3
	// nums1 := []int{}
	// m := 0
	// nums2 := []int{1}
	// n := 1

	//case 4
	// nums1 := []int{1, 0}
	// m := 1
	// nums2 := []int{2}
	// n := 1

	//case 5
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)

	fmt.Println(" Result = ", nums1)
}
