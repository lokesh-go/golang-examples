package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	return inOrder(root, &k)
}

func inOrder(root *TreeNode, k *int) int {
	if root == nil {
		return -1
	}
	leftRes := inOrder(root.Left, k)
	if leftRes != -1 {
		return leftRes
	}
	*k--
	if *k == 0 {
		return root.Val
	}
	rightRes := inOrder(root.Right, k)
	if rightRes != -1 {
		return rightRes
	}
	return -1
}
