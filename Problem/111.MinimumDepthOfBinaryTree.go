package Problem

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := minDepth(root.Left), minDepth(root.Right)
	if l == 0 || r == 0 {
		return l + r + 1
	}
	return getMin(l, r) + 1
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-21 08:38:56

//Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the roo
//t node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Given binary tree [3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// return its minimum depth = 2.
//
// 👍 321 👎 0
