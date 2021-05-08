package main

func swapPairs(head *ListNode) *ListNode {
	num := 0
	var parent *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	c := head.Next
	gc := head.Next.Next
	head = c
	head.Next = p
	head.Next.Next = gc
	grandParent := head.Next
	node := head.Next.Next
	for {
		if node == nil {
			if parent != nil {

			}
			break
		}
		if num%2 == 0 {
			parent = node
		} else {
			child := *node
			gc := child.Next
			grandParent.Next = &child
			grandParent.Next.Next = parent
			grandParent.Next.Next.Next = gc
			grandParent = parent
			parent = nil
		}
		node = node.Next
		num++
	}
	return head
}
