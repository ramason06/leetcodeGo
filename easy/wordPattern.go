package easy

import "strings"

func wordPattern(pattern string, s string) bool {
	m := make(map[rune]string)
	tm := make(map[string]rune)

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i, r := range pattern {
		sc := r
		wc := words[i]
		if val, ex := m[sc]; ex {
			if val != wc {
				return false
			}
		} else {
			m[sc] = wc
		}

		if val, ex := tm[wc]; ex {
			if val != sc {
				return false
			}
		} else {
			tm[wc] = sc
		}
	}

	return true
}
