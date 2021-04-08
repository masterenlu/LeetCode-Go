package Problem

//
// Given a binary tree, you need to compute the length of the diameter of the tre
// e. The diameter of a binary tree is the length of the longest path between any t
// wo nodes in a tree. This path may or may not pass through the root.
//
//
//
// Example:
// Given a binary tree
//
//          1
//         / \
//        2   3
//       / \
//      4   5
//
//
//
// Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
//
//
// Note:
// The length of path between two nodes is represented by the number of edges bet
// ween them.
// Related Topics 树
// 👍 420 👎 0

/*
	测算每个结点深度，结点的宽度就是两个子结点的最大深度之和
*/

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maxWidth int

//goland:noinspection GoUnusedFunction
func diameterOfBinaryTree(root *TreeNode) int {
	maxWidth = 0
	maxDepth(root)
	return maxWidth
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lD, rD := maxDepth(root.Left), maxDepth(root.Right)
	maxWidth = max(maxWidth, lD+rD)
	return max(lD, rD) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-07-28 09:18:11
