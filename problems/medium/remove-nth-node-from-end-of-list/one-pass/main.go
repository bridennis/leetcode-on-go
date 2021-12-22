package medium_remove_nth_node_from_end_of_list_one_pass

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/
func removeNthFromEndIn(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	p2 := head
	for ; n > 0; n-- {
		p2 = p2.Next
	}

	if p2 == nil {
		return head.Next
	}

	p1 := head
	for p2.Next != nil {
		p1, p2 = p1.Next, p2.Next
	}

	p1.Next = p1.Next.Next

	return head
}
