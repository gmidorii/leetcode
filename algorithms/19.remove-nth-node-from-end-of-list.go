package main

import "fmt"

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

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var fast, slow *ListNode
	fast = head
	slow = head
	step := 0
	for i := 0; i < n; i++ {
		if fast.Next == nil && step < n-1 {
			return head
		}
		fmt.Printf("%v: %v\n", i, fast.Val)
		fmt.Printf("%v: %v\n", i, slow.Val)
		fast = fast.Next
		step++
	}

	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
