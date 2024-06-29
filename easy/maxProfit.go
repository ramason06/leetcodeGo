package easy

func maxProfit(prices []int) int {
	s := 0
	p := 0
	for _, val := range prices {
		if s == 0 || val < s {
			s = val
		} else if val-s > p {
			p = val - s
		}
	}

	return p
}
