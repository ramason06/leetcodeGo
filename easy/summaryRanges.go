package easy

import "strconv"

func summaryRanges(nums []int) []string {
	var ans []string

	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	a, s := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if i+1 == len(nums) {
			if a+1 == nums[i] {
				ans = append(ans, strconv.Itoa(s)+"->"+strconv.Itoa(a+1))
			} else {
				if s == a {
					ans = append(ans, strconv.Itoa(s))
				} else {
					ans = append(ans, strconv.Itoa(s)+"->"+strconv.Itoa(a))
				}
				ans = append(ans, strconv.Itoa(nums[i]))
			}
		} else if a+1 == nums[i] {
			a = nums[i]
		} else {
			if s == a {
				ans = append(ans, strconv.Itoa(s))
			} else {
				ans = append(ans, strconv.Itoa(s)+"->"+strconv.Itoa(a))
			}
			s = nums[i]
			a = nums[i]
		}
	}
	return ans
}

func summaryRangesSimple(nums []int) []string {
	var ans []string
	n := len(nums)
	if n == 0 {
		return ans
	}

	s := nums[0]
	for i := 1; i <= n; i++ {
		if i == n || nums[i] != nums[i-1]+1 {
			if nums[i-1] == s {
				ans = append(ans, strconv.Itoa(s))
			} else {
				ans = append(ans, strconv.Itoa(s)+"->"+strconv.Itoa(nums[i-1]))
			}
			if i < n {
				s = nums[i]
			}
		}
	}

	return ans
}
