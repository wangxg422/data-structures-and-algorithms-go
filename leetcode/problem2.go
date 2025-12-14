package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// from AI
func addTwoNumbersAI(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
	}

	return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := r

	m := 0 // 进位标志
	for l1 != nil || l2 != nil {
		sum := m
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		if sum > 9 {
			tail.Next = &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			m = 1
		} else {
			tail.Next = &ListNode{
				Val:  sum,
				Next: nil,
			}
			m = 0
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		tail = tail.Next
	}

	if m > 0 {
		tail.Next = &ListNode{
			Val:  m,
			Next: nil,
		}
	}

	return r.Next
}
