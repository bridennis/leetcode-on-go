package medium_populating_next_right_pointers_in_each_node

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Connect #1",
			args: struct{ root *Node }{root: tree([]int{1, 2, 3, 4, 5, 6, 7})},
			want: connectLevels(tree([]int{1, 2, 3, 4, 5, 6, 7})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func connectLevels(root *Node) *Node {
	root.Left.Next = root.Right
	root.Left.Left.Next = root.Left.Right
	root.Left.Right.Next = root.Right.Left
	root.Right.Left.Next = root.Right.Right

	return root
}

func tree(v []int) *Node {
	if len(v) == 0 {
		return nil
	}

	head := &Node{}
	head = buildTree(v, head, 0)

	return head
}

func buildTree(v []int, head *Node, i int) *Node {
	if i >= len(v) || v[i] == -1 {
		return head
	}

	node := &Node{
		Val:   v[i],
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	node.Left = buildTree(v, node.Left, 2*i+1)
	node.Right = buildTree(v, node.Right, 2*i+2)

	return node
}
