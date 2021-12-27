package medium_populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
*/
func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	populate(root.Left, root)
	populate(root.Right, root)

	return root
}

func populate(n *Node, parent *Node) {
	if parent.Right == n && parent.Next != nil {
		n.Next = parent.Next.Left
	} else if parent.Left == n {
		n.Next = parent.Right
	}

	if n.Left != nil {
		populate(n.Left, n)
	}
	if n.Right != nil {
		populate(n.Right, n)
	}
}
