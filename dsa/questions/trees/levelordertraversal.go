package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func levelOrderTraversal(root *TreeNode) (res [][]int) {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var levelNodes []int
		for i := 0; i < levelSize; i++ {
			// Pop an element
			node := queue[0]
			queue = queue[1:]

			levelNodes = append(levelNodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelNodes)
	}
	return res
}

/*
func main() {
	// Example usage:
	// Construct the binary tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	// Perform level order traversal
	result := levelOrderTraversal(root)

	// Print the result
	for _, level := range result {
		fmt.Println(level)
	}
}
*/
