package Problem

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//goland:noinspection GoUnusedFunction
func recoverTree(root *TreeNode) {
	var x, y, pre, predecessor *TreeNode

	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				if pre != nil && pre.Val > root.Val {
					y = root
					if x == nil {
						x = pre
					}
				}
				pre = root
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			if pre != nil && pre.Val > root.Val {
				y = root
				if x == nil {
					x = pre
				}
			}
			pre = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-08 09:09:01

// Two elements of a binary search tree (BST) are swapped by mistake.
//
// Recover the tree without changing its structure.
//
// Example 1:
//
//
// Input: [1,3,null,null,2]
//
//    1
//   /
//  3
//   \
//    2
//
// Output: [3,1,null,null,2]
//
//    3
//   /
//  1
//   \
//    2
//
//
// Example 2:
//
//
// Input: [3,1,4,null,null,2]
//
//  3
// / \
// 1   4
//    /
//   2
//
// Output: [2,1,4,null,null,3]
//
//  2
// / \
// 1   4
//    /
//  3
//
//
// Follow up:
//
//
// A solution using O(n) space is pretty straight forward.
// Could you devise a constant space solution?
//
//
// 👍 271 👎 0
