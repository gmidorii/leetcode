package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	var l []*ListNode
	if node.Next == nil {
		return nil
	} else {
		l = append(l, node)
	}
	for {
		if node.Next == nil {
			if len(l)-n-1 < 0 {
				return l[1]
			}
			if len(l)-n+1 >= len(l) {
				l[len(l)-2].Next = nil
				return head
			}
			parent := l[len(l)-n-1]
			grandChild := l[len(l)-n+1]
			parent.Next = grandChild
			break
		}
		l = append(l, node.Next)
		node = node.Next
	}
	return head
}
