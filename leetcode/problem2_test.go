package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println("l1: " + printNode(l1))
	fmt.Println("l2: " + printNode(l2))

	r := addTwoNumbers(l1, l2)
	fmt.Println("result: " + printNode(r))
}

func printNode(l *ListNode) string {
	s := make([]string, 0)
	for l != nil {
		s = append(s, fmt.Sprintf("%d", l.Val))
		l = l.Next
	}

	return strings.Join(s, "->")
}
