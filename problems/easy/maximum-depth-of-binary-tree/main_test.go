package easy_maximum_depth_of_binary_tree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maxDepth #1",
			args: struct{ root *TreeNode }{root: tree([]int{3, 9, 20, -1, -1, 15, 7})},
			want: 3,
		},
		{
			name: "maxDepth #2",
			args: struct{ root *TreeNode }{root: tree([]int{1, -1, 2})},
			want: 2,
		},
		{
			name: "maxDepth #3",
			args: struct{ root *TreeNode }{root: tree([]int{})},
			want: 0,
		},
		{
			name: "maxDepth #4",
			args: struct{ root *TreeNode }{root: tree([]int{1})},
			want: 1,
		},
		{
			name: "maxDepth #5",
			args: struct{ root *TreeNode }{root: tree([]int{1, 2})},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
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
