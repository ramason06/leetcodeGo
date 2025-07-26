package easy

import "strconv"

func isPalindromeInt(x int) bool {
	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < len(s)/2; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
