package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Accepted
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	slow := dummy
	fast := dummy

	step := 0
	for fast != nil {
		if step <= n {
			fast = fast.Next
			step++
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return dummy.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{0, head}
	fast, slow := dummyNode, dummyNode
	for i := 0; i <= n; i++ { // 注意<=，否则快指针为空时，慢指针正好在倒数第n个上面
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	newHead := removeNthFromEnd(head, 2)

	for newHead != nil {
		fmt.Printf("%d ", newHead.Val)
		newHead = newHead.Next
	}
	fmt.Println()
}
