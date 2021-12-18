package medium_add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Summarize two lists #1",
			args: struct {
				l1 *ListNode
				l2 *ListNode
			}{
				l1: list([]int{2, 4, 3}),
				l2: list([]int{5, 6, 4}),
			},
			want: list([]int{7, 0, 8}),
		},
		{
			name: "Summarize two lists #2",
			args: struct {
				l1 *ListNode
				l2 *ListNode
			}{
				l1: list([]int{9}),
				l2: list([]int{9}),
			},
			want: list([]int{8, 1}),
		},
		{
			name: "Summarize two lists #3",
			args: struct {
				l1 *ListNode
				l2 *ListNode
			}{
				l1: list([]int{2, 4, 9}),
				l2: list([]int{5, 6, 4, 9}),
			},
			want: list([]int{7, 0, 4, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
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
