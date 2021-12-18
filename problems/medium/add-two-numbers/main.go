package medium_add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/add-two-numbers/
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}

	if l1 != nil && l2 == nil {
		return l1
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := head
	var rest int

	for {
		if l1 == nil && l2 == nil {
			break
		}

		current.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		current = current.Next

		if l1 != nil && l2 != nil {
			current.Val = (l1.Val + l2.Val + rest) % 10
			rest = (l1.Val + l2.Val + rest) / 10

			l1 = l1.Next
			l2 = l2.Next

			continue
		} else if l1 == nil {
			l1, l2 = l2, l1
		}

		current.Val = (l1.Val + rest) % 10
		rest = (l1.Val + rest) / 10

		l1 = l1.Next
	}

	if rest > 0 {
		current.Next = &ListNode{
			Val:  rest,
			Next: nil,
		}
	}

	return head.Next
}
