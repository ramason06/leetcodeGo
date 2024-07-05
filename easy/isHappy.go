package easy

func isHappy(n int) bool {
	mapp := make(map[int]bool)
	for n != 1 {
		pre := n
		sum := 0
		for pre > 0 {
			m := pre % 10
			sum += m * m
			pre /= 10
		}
		n = sum
		if _, ex := mapp[sum]; ex {
			return false
		}
		mapp[sum] = true
	}
	return true
}
