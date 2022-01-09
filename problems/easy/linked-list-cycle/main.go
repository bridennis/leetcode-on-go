package easy_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/linked-list-cycle/
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p1 := head
	p2 := head.Next
	for {
		if p2 == nil {
			return false
		}

		if p2 == p1 {
			return true
		}

		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		}
	}
}
