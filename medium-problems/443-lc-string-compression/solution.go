package lcstringcompression

import "fmt"

func compress(chars []byte) int {
	count := 0
	lastIndex := 0
	for i := 0; i < len(chars); i++ {
		ch := chars[i]
		count = 0

		for i < len(chars) && ch == chars[i] {
			count++
			i++
		}
		if count == 1 {
			chars[lastIndex] = ch
			lastIndex++
		} else {
			chars[lastIndex] = ch
			lastIndex++
			// convert count to string and add to chars
			countStr := fmt.Sprintf("%d", count)
			for _, v := range countStr {
				chars[lastIndex] = byte(v)
				lastIndex++
			}
		}
		i--
	}
	return lastIndex
}
