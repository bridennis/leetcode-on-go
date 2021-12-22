package medium_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	node := head
	nodes := 1
	for ; node.Next != nil; nodes++ {
		node = node.Next
	}

	node = head
	for skip := nodes - n; node.Next != nil; skip-- {
		if skip == 0 {
			head = head.Next
			break
		}

		if skip == 1 {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
			} else {
				node.Next = nil
			}
			break
		}
		node = node.Next
	}

	return head
}
