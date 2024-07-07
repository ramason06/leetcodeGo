package easy

import (
	"leetcodeGo/common"
)

func isValid(s string) bool {
	sk1 := &common.Stack{}

	for _, val := range s {
		var e error
		var l string
		if val == '(' {
			sk1.Push(string(val))
		} else if val == ')' {
			l, e = sk1.Pop()
			if l != "(" {
				return false
			}
		} else if val == '{' {
			sk1.Push(string(val))
		} else if val == '}' {
			l, e = sk1.Pop()
			if l != "{" {
				return false
			}
		} else if val == '[' {
			sk1.Push(string(val))
		} else {
			l, e = sk1.Pop()
			if l != "[" {
				return false
			}
		}

		if e != nil {
			return false
		}
	}
	if sk1.Size() == 0 {
		return true
	}
	return false
}
