package trees

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return isBST(root, math.MinInt, math.MaxInt)
}

func isBST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	if node.Val >= max || node.Val <= min {
		return false
	}
	return isBST(node.Left, min, node.Val) && isBST(node.Right, node.Val, max)
}
