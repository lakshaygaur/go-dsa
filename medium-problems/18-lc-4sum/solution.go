package lc4sum

import (
	"encoding/json"
	"sort"
)

func isDistinct(a int, b int, c int, d int) bool {
	if a != b && b != c && b != d && c != a && c != d && a != d {
		return true
	} else {
		return false
	}
}

func fourSum(nums []int, target int) [][]int {
	numsMap := make(map[int]int)
	var answers [][]int
	ansMap := make(map[string][]int)

	for i, v := range nums {
		numsMap[v] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				desired := target - nums[i] - nums[j] - nums[k]
				if _, ok := numsMap[desired]; ok && isDistinct(numsMap[desired], i, j, k) {
					// add that to answers
					ans := []int{desired, nums[i], nums[j], nums[k]}
					sort.Ints(ans)
					// answers=append(answers, ans)
					s, _ := json.Marshal(ans)
					ansMap[string(s)] = ans
				}
			}
		}
	}

	for _, v := range ansMap {
		answers = append(answers, v)
	}
	return answers
}
