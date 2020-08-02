package Problem

//Given a binary tree, flatten it to a linked list in-place.
//
// For example, given the following tree:
//
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//
//
// The flattened tree should look like:
//
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
//
// Related Topics 树 深度优先搜索
// 👍 441 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	p := root
	for p != nil {
		if p.Left != nil {
			temp := p.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			temp.Right, p.Right, p.Left = p.Right, p.Left, nil
		}
		p = p.Right
	}
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-02 08:21:41
