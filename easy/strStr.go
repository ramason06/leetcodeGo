package easy

import "strings"

func indexOf(s, substr string) int {
	n, m := len(s), len(substr)
	if m == 0 {
		return 0
	}
	if n < m {
		return -1
	}

	// 부분 문자열을 찾기 위한 루프
	for i := 0; i <= n-m; i++ {
		if s[i:i+m] == substr {
			return i
		}
	}

	return -1
}

// Find the Index of the First Occurrence in a String
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
