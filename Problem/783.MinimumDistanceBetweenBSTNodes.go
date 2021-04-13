package Problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	ans := 105
	for len(nodes) != 0 && ans != 0 {
		r := nodes[0]
		nodes = nodes[1:]
		if r.Left != nil {
			nodes = append(nodes, r.Left)
			leftVal := getLeftMaxValue(r.Left)
			ans = getMin(r.Val-leftVal, ans)
		}
		if r.Right != nil {
			nodes = append(nodes, r.Right)
			rightVal := getRightMaxValue(r.Right)
			ans = getMin(rightVal-r.Val, ans)
		}
	}
	return ans
}

func getLeftMaxValue(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return getLeftMaxValue(root.Right)
}

func getRightMaxValue(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return getRightMaxValue(root.Left)
}

func getMin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-13 08:44:44

// Given the root of a Binary Search Tree (BST), return the minimum difference be
// tween the values of any two different nodes in the tree.
//
// Note: This question is the same as 530: https://leetcode.com/problems/minimum
// -absolute-difference-in-bst/
//
//
// Example 1:
//
//
// Input: root = [4,2,6,1,3]
// Output: 1
//
//
// Example 2:
//
//
// Input: root = [1,0,48,null,null,12,49]
// Output: 1
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [2, 100].
// 0 <= Node.val <= 105
//
//
// 👍 127 👎 0
