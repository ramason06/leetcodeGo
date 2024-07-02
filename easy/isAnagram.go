package easy

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for i, val := range s {
		m[val]++
		m[rune(t[i])]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
