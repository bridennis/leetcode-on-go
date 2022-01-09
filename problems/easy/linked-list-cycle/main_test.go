package easy_linked_list_cycle

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Has cycle #1",
			args: struct{ head *ListNode }{head: cycledAt(list([]int{3, 2, 0, -4}), 1)},
			want: true,
		},
		{
			name: "Has cycle #2",
			args: struct{ head *ListNode }{head: cycledAt(list([]int{1, 2}), 0)},
			want: true,
		},
		{
			name: "Has cycle #3",
			args: struct{ head *ListNode }{head: cycledAt(list([]int{1}), -1)},
			want: false,
		},
		{
			name: "Has cycle #4",
			args: struct{ head *ListNode }{head: cycledAt(list([]int{1, 2, 3}), -1)},
			want: false,
		},
		{
			name: "Has cycle #5",
			args: struct{ head *ListNode }{head: cycledAt(list([]int{1, 2}), -1)},
			want: false,
		},
		{
			name: "Has cycle #6",
			args: struct{ head *ListNode }{head: cycledAt(list([]int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5}), 6)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func cycledAt(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil || n < 0 {
		return head
	}

	pos := 0
	var node *ListNode
	curNode := head
	for curNode.Next != nil {
		if pos == n {
			node = curNode
		}
		curNode = curNode.Next
		pos++
	}

	curNode.Next = node

	return head
}

func list(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  l[0],
		Next: nil,
	}
	current := head
	for i := 1; i < len(l); i++ {
		current.Next = &ListNode{
			Val:  l[i],
			Next: nil,
		}
		current = current.Next
	}

	return head
}
