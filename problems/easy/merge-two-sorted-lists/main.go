package easy_merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}

	if list2 == nil && list1 != nil {
		return list1
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := head

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				list1, list2 = list2, list1
			}
		} else if list2 != nil {
			list1, list2 = list2, list1
		}

		current.Next, current = list1, list1
		list1 = list1.Next
	}

	return head.Next
}
