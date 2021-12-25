package easy_merge_two_binary_trees

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "Merge #1", args: struct {
			root1 *TreeNode
			root2 *TreeNode
		}{
			root1: tree([]int{1, 3, 2, 5}),
			root2: tree([]int{2, 1, 3, -1, 4, -1, 7}),
		}, want: tree([]int{3, 4, 5, 5, 4, -1, 7})},
		{name: "Merge #2", args: struct {
			root1 *TreeNode
			root2 *TreeNode
		}{
			root1: tree([]int{1}),
			root2: tree([]int{1, 2}),
		}, want: tree([]int{2, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func tree(v []int) *TreeNode {
	if len(v) == 0 {
		return nil
	}

	head := &TreeNode{}
	head = buildTree(v, head, 0)

	return head
}

func buildTree(v []int, head *TreeNode, i int) *TreeNode {
	if i >= len(v) || v[i] == -1 {
		return head
	}

	node := &TreeNode{
		Val:   v[i],
		Left:  nil,
		Right: nil,
	}
	node.Left = buildTree(v, node.Left, 2*i+1)
	node.Right = buildTree(v, node.Right, 2*i+2)

	return node
}
