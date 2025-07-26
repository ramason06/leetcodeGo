package easy

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carry == 1 {
			digits[i]++
			if (digits[i] / 10) > 0 {
				carry = 1
			} else {
				carry = 0
			}
			digits[i] %= 10
		}
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
