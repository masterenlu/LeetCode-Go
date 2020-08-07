package Problem

//Given two binary trees, write a function to check if they are the same or not.
//
//
// Two binary trees are considered the same if they are structurally identical a
//nd the nodes have the same value.
//
// Example 1:
//
//
//Input:     1         1
//          / \       / \
//         2   3     2   3
//
//        [1,2,3],   [1,2,3]
//
//Output: true
//
//
// Example 2:
//
//
//Input:     1         1
//          /           \
//         2             2
//
//        [1,2],     [1,null,2]
//
//Output: false
//
//
// Example 3:
//
//
//Input:     1         1
//          / \       / \
//         2   1     1   2
//
//        [1,2,1],   [1,1,2]
//
//Output: false
//
// Related Topics 树 深度优先搜索
// 👍 418 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-07 09:19:30
