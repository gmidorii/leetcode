package main

func swapPairs(head *ListNode) *ListNode {
	num := 0
	var grandParent *ListNode
	var parent *ListNode
	var child *ListNode
	node := head
	for {
		if node == nil {
			break
		}
		if num%2 == 1 {
			parent = node
			node = node.Next
			continue
		} else {
			child = node
			if grandParent == nil {
				grandChild := child.Next
				child.Next = parent
				parent.Next = grandChild
				grandParent = child
			} else {
				grandChild := child.Next
				child.Next = parent
				grandParent.Next = child
				parent.Next = grandChild
			}
			node = node.Next
		}
	}
	return head
}
