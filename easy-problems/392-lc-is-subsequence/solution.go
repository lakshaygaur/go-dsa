package lcissubsequence

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	count := 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			count++
			i++
			j++
		} else { // characters dont match
			j++
		}
	}

	if count == len(s) {
		return true
	}
	return false
}
