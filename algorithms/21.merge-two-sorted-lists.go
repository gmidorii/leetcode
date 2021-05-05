package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var node *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			if head == nil {
				head = l2
				node = head
				l2 = l2.Next
				continue
			}
			node.Next = l2
			node = node.Next
			l2 = l2.Next
			continue
		} else if l2 == nil {
			if head == nil {
				head = l1
				node = head
				l1 = l1.Next
				continue
			}
			node.Next = l1
			node = node.Next
			l1 = l1.Next
			continue
		} else {
			if head == nil {
				if l1.Val > l2.Val {
					head = l2
					node = head
					l2 = l2.Next
				} else {
					head = l1
					node = head
					l1 = l1.Next
				}
				continue
			}
			if l1.Val > l2.Val {
				node.Next = l2
				node = node.Next
				l2 = l2.Next
			} else {
				node.Next = l1
				node = node.Next
				l1 = l1.Next
			}
			continue
		}
	}
	return head
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoLists2(l1.Next, l2)
	return l1
}
