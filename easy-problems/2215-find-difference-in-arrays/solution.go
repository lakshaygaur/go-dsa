package find_diff_in_arr

func findDifference(nums1 []int, nums2 []int) [][]int {

	outMap1 := make(map[int]int)
	outMap2 := make(map[int]int)
	var output [][]int
	var outNums1 []int
	var outNums2 []int

	for _, e := range nums1 {
		outMap1[e]++
	}

	for _, e := range nums2 {
		outMap2[e]++
	}

	for key, _ := range outMap1 {
		if outMap2[key] == 0 {
			outNums1 = append(outNums1, key)
		}
	}

	for e, _ := range outMap2 {
		if outMap1[e] == 0 {
			outNums2 = append(outNums2, e)
		}
	}

	output = append(output, outNums1)
	output = append(output, outNums2)

	return output
}
