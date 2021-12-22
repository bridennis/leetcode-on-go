package easy_middle_of_the_linked_list

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "MiddleNode #1",
			args: struct{ head *ListNode }{head: list([]int{1, 2, 3, 4, 5})},
			want: list([]int{3, 4, 5}),
		},
		{
			name: "MiddleNode #2",
			args: struct{ head *ListNode }{head: list([]int{1, 2, 3, 4, 5, 6})},
			want: list([]int{4, 5, 6}),
		},
		{
			name: "MiddleNode #2",
			args: struct{ head *ListNode }{head: list([]int{1})},
			want: list([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
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
