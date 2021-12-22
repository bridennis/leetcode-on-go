package medium_remove_nth_node_from_end_of_list_one_pass

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEndIn(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "RemoveNthFromEnd #1",
			args: struct {
				head *ListNode
				n    int
			}{head: list([]int{1, 2, 3, 4, 5}), n: 2},
			want: list([]int{1, 2, 3, 5}),
		},
		{
			name: "RemoveNthFromEnd #2",
			args: struct {
				head *ListNode
				n    int
			}{head: list([]int{1}), n: 1},
			want: list([]int{}),
		},
		{
			name: "RemoveNthFromEnd #3",
			args: struct {
				head *ListNode
				n    int
			}{head: list([]int{1, 2}), n: 1},
			want: list([]int{1}),
		},
		{
			name: "RemoveNthFromEnd #4",
			args: struct {
				head *ListNode
				n    int
			}{head: list([]int{1, 2}), n: 2},
			want: list([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEndIn(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEndIn() = %v, want %v", got, tt.want)
			}
		})
	}
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
