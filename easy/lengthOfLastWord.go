package easy

func lengthOfLastWord(s string) int {
	cnt, maxx := 0, 0
	for _, c := range s {
		if c != ' ' {
			cnt++
			maxx = cnt
		} else if c == ' ' {
			cnt = 0
		}
	}
	return maxx
}
