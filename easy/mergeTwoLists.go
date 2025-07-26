package easy

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 != nil && (list2 == nil || (list2 != nil && list1.Val <= list2.Val)) {
		head = list1
		list1 = list1.Next
	} else if list2 != nil {
		head = list2
		list2 = list2.Next
	}
	var temp *ListNode
	temp = head
	for list1 != nil || list2 != nil {
		if list1 != nil && (list2 == nil || list1.Val <= list2.Val) {
			head.Next = list1
			list1 = list1.Next
		} else if list2 != nil {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	return temp
}
