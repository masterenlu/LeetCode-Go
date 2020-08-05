package Problem

//The thief has found himself a new place for his thievery again. There is only
//one entrance to this area, called the "root." Besides the root, each house has o
//ne and only one parent house. After a tour, the smart thief realized that "all h
//ouses in this place forms a binary tree". It will automatically contact the poli
//ce if two directly-linked houses were broken into on the same night.
//
// Determine the maximum amount of money the thief can rob tonight without alert
//ing the police.
//
// Example 1:
//
//
//Input: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//Output: 7
//Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
//
// Example 2:
//
//
//Input: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//Output: 9
//Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
// Related Topics 树 深度优先搜索
// 👍 454 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	ans1, ans2 := robDP(root)
	return getMax(ans1, ans2)
}

func robDP(root *TreeNode) (val1, val2 int) {
	if root == nil {
		return 0, 0
	}

	left1, left2 := robDP(root.Left)
	right1, right2 := robDP(root.Right)
	val1 = left2 + right2 + root.Val
	val2 = getMax(left1, left2) + getMax(right1, right2)
	return
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-05 09:11:27
