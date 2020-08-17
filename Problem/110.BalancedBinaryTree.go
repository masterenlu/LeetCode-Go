package Problem

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	flag := true
	_, flag = dfs(root, flag)
	return flag
}

func dfs(root *TreeNode, flag bool) (int, bool) {
	if root == nil {
		return 0, flag
	}
	l, flag1 := dfs(root.Left, flag)
	r, flag2 := dfs(root.Right, flag)
	if !flag1 || !flag2 || math.Abs(float64(l-r)) > 1.0 {
		return 0, false
	}
	return getMax(l, r) + 1, flag
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-17 09:03:08

//Given a binary tree, determine if it is height-balanced.
//
// For this problem, a height-balanced binary tree is defined as:
//
//
// a binary tree in which the left and right subtrees of every node differ in he
//ight by no more than 1.
//
//
//
//
// Example 1:
//
// Given the following tree [3,9,20,null,null,15,7]:
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// Return true.
//
//Example 2:
//
// Given the following tree [1,2,2,3,3,null,null,4,4]:
//
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//
//
// Return false.
//
// 👍 408 👎 0
