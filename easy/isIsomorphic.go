package easy

func isIsomorphic(s string, t string) bool {
	m := make(map[rune]rune)
	tm := make(map[rune]rune)

	for i, r := range s {
		sc := r
		tc := rune(t[i])
		if val, ex := m[sc]; ex {
			if val != tc {
				return false
			}
		} else {
			m[sc] = tc
		}

		if val, ex := tm[tc]; ex {
			if val != sc {
				return false
			}
		} else {
			tm[tc] = sc
		}
	}

	return true
}
