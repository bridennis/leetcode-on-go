package easy_maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
https://leetcode.com/problems/maximum-depth-of-binary-tree/
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	queue := []*TreeNode{root, nil}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode == nil {
			maxDepth++
		}

		if curNode != nil {
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		} else if len(queue) > 0 {
			queue = append(queue, nil)
		}
	}

	return maxDepth
}
