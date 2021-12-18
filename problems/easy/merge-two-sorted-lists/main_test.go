package easy_merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Merge two lists",
			args: struct {
				list1 *ListNode
				list2 *ListNode
			}{
				list1: list([]int{1, 2, 4}),
				list2: list([]int{1, 3, 4}),
			},
			want: list([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
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
