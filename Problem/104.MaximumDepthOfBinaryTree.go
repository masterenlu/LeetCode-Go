package Problem

//Given a binary tree, find its maximum depth.
//
// The maximum depth is the number of nodes along the longest path from the root
// node down to the farthest leaf node.
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
// return its depth = 3.
// Related Topics 树 深度优先搜索
// 👍 618 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-28 08:54:58
