package easy_reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/reverse-linked-list/
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, next *ListNode
	current := head
	for current != nil {
		next, current.Next = current.Next, prev
		prev, current = current, next
	}

	return prev
}
