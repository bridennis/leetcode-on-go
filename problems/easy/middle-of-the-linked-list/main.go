package easy_middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/middle-of-the-linked-list/
*/
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodes := 1
	node := head
	for node.Next != nil {
		nodes++
		node = node.Next
	}

	node = head
	for jumpTo := nodes / 2; jumpTo > 0; jumpTo-- {
		node = node.Next
	}

	return node
}
