package uniqueoccurences

func uniqueOccurrences(arr []int) bool {
	nMap := make(map[int]int)
	nMapValue := make(map[int][]int)
	for _, v := range arr {
		nMap[v]++
	}

	for k, v := range nMap {
		nMapValue[v] = append(nMapValue[v], k)
	}

	for _, v := range nMapValue {
		if len(v) > 1 {
			return false
		}
	}

	return true
}
