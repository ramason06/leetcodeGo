package easy

import "unicode"

func isPalindrome(s string) bool {
	alpha := make([]rune, 0, len(s)) // 결과 문자열을 저장할 슬라이스 초기화

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			alpha = append(alpha, unicode.ToLower(char))
		}
	}

	l := len(alpha)
	i, j := 0, l-1

	for i < l/2 && j >= l/2 {
		if alpha[i] != alpha[j] {
			return false
		}
		i++
		j--
	}

	return true
}
