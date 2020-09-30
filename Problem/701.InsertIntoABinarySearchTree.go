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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	temp := root
	for {
		if temp.Val > val {
			if temp.Left != nil {
				temp = temp.Left
				continue
			}
			temp.Left = &TreeNode{Val: val}
			break
		}
		if temp.Val < val {
			if temp.Right != nil {
				temp = temp.Right
				continue
			}
			temp.Right = &TreeNode{Val: val}
			break
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-09-30 23:04:07

//Given the root node of a binary search tree (BST) and a value to be inserted i
//nto the tree, insert the value into the BST. Return the root node of the BST aft
//er the insertion. It is guaranteed that the new value does not exist in the orig
//inal BST.
//
// Note that there may exist multiple valid ways for the insertion, as long as t
//he tree remains a BST after insertion. You can return any of them.
//
// For example,
//
//
//Given the tree:
//        4
//       / \
//      2   7
//     / \
//    1   3
//And the value to insert: 5
//
//
// You can return this binary search tree:
//
//
//         4
//       /   \
//      2     7
//     / \   /
//    1   3 5
//
//
// This tree is also valid:
//
//
//         5
//       /   \
//      2     7
//     / \
//    1   3
//         \
//          4
//
//
//
// Constraints:
//
//
// The number of nodes in the given tree will be between 0 and 10^4.
// Each node will have a unique integer value from 0 to -10^8, inclusive.
// -10^8 <= val <= 10^8
// It's guaranteed that val does not exist in the original BST.
//
//
// 👍 126 👎 0
