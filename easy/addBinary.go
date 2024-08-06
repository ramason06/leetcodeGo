package easy

func addBinary(a string, b string) string {
	al, bl := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for al >= 0 || bl >= 0 || carry != 0 {
		var av, bv int

		if al >= 0 {
			av = int(a[al] - '0')
			al--
		} else {
			av = 0
		}

		if bl >= 0 {
			bv = int(b[bl] - '0')
			bl--
		} else {
			bv = 0
		}

		sum := av + bv + carry
		result = string(sum%2+'0') + result
		carry = sum / 2
	}

	return result
}
