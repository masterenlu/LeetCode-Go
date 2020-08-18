package Problem

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return createTree(head, nil)
}

func getMid(left, right *ListNode) *ListNode {
	fast, slow := left, left

	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func createTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMid(left, right)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = createTree(left, mid)
	root.Right = createTree(mid.Next, right)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-18 08:56:12

//Given the head of a singly linked list where elements are sorted in ascending
//order, convert it to a height balanced BST.
//
// For this problem, a height-balanced binary tree is defined as a binary tree i
//n which the depth of the two subtrees of every node never differ by more than 1.
//
//
//
// Example 1:
//
//
//Input: head = [-10,-3,0,5,9]
//Output: [0,-3,9,-10,null,5]
//Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the
//shown height balanced BST.
//
//
// Example 2:
//
//
//Input: head = []
//Output: []
//
//
// Example 3:
//
//
//Input: head = [0]
//Output: [0]
//
//
// Example 4:
//
//
//Input: head = [1,3]
//Output: [3,1]
//
//
//
// Constraints:
//
//
// The numner of nodes in head is in the range [0, 2 * 10^4].
// -10^5 <= Node.val <= 10^5
//
//
// 👍 290 👎 0
