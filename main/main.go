package main

func main() {
	//a := new(ListNode)
	//a.Val = 3
	//b := new(ListNode)
	//b.Val = 2
	//b.Next = a
	c := new(ListNode)
	c.Val = 1
	c.Next = nil
	//
	//d := new(ListNode)
	//d.Val = 4
	//f := new(ListNode)
	//f.Val = 3
	//f.Next = d
	//e := new(ListNode)
	//e.Val = 1
	//e.Next = f
	mergeTwoLists(nil, c)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

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
