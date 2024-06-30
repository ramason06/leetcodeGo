package easy

func canConstruct(ransomNote string, magazine string) bool {
	var m = make(map[rune]int)

	for _, val := range ransomNote {
		m[val]++
	}

	for _, val := range magazine {
		m[val]--
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
