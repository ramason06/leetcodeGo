package easy

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a := 0
	b := head
	for b != nil {
		a++
		b = b.Next
	}
	if a == 1 && n == 1 {
		return nil
	}

	position := a - n
	if n != 0 {
		if position > 0 {
			sol := head
			for i := 0; i <= position-1; i++ {
				sol = sol.Next
			}
			sol.Next = sol.Next.Next
		} else if a == n {
			head = head.Next
		}
	}
	return head
}
