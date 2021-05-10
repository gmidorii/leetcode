package main

import "fmt"

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

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s := head.Next
	var behind *ListNode
	for head.Next != nil {
		headNext := head.Next
		if behind != nil && behind.Next != nil {
			behind.Next = headNext
		}
		var headNextNext *ListNode
		if head.Next.Next != nil {
			headNextNext = head.Next.Next
		}
		head.Next = headNextNext
		headNext.Next = head
		behind = head
		if head.Next != nil {
			head = headNextNext
		}
	}
	return s
}

func printListNode(n *ListNode) {
	for n != nil {
		fmt.Printf("%v, ", n.Val)
		n = n.Next
	}
	fmt.Println("")
}
