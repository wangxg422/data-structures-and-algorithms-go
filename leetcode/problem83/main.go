package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head

	for p != nil {
		if p.Next != nil && p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
		},
	}

	r := deleteDuplicates(head)

	for r != nil {
		fmt.Printf("%d -> ", r.Val)
		r = r.Next
	}
}
