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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1, arr2 := getArray(root1), getArray(root2)
	ans := true
	if len(arr1) != len(arr2) {
		ans = false
	} else {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] != arr2[i] {
				ans = false
				break
			}
		}
	}
	return ans
}

func getArray(root *TreeNode) []int {
	var arr []int
	list := []*TreeNode{root, root}
	for len(list) != 1 {
		p := list[len(list)-1]
		list = list[:len(list)-1]
		if p.Right != nil {
			list = append(list, p.Right)
		}
		if p.Left != nil {
			list = append(list, p.Left)
		} else if p.Right == nil {
			arr = append(arr, p.Val)
		}
	}
	return arr
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-10 09:45:53

// Consider all the leaves of a binary tree, from left to right order, the values
// of those leaves form a leaf value sequence.
//
//
//
// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9,
// 8).
//
// Two binary trees are considered leaf-similar if their leaf value sequence is
// the same.
//
// Return true if and only if the two given trees with head nodes root1 and root
// 2 are leaf-similar.
//
//
// Example 1:
//
//
// Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null
// ,null,null,null,null,9,8]
// Output: true
//
//
// Example 2:
//
//
// Input: root1 = [1], root2 = [1]
// Output: true
//
//
// Example 3:
//
//
// Input: root1 = [1], root2 = [2]
// Output: false
//
//
// Example 4:
//
//
// Input: root1 = [1,2], root2 = [2,2]
// Output: true
//
//
// Example 5:
//
//
// Input: root1 = [1,2,3], root2 = [1,3,2]
// Output: false
//
//
//
// Constraints:
//
//
// The number of nodes in each tree will be in the range [1, 200].
// Both of the given trees will have values in the range [0, 200].
//
//
// 👍 115 👎 0
