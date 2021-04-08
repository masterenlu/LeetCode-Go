package 面试题

// A magic index in an array A[0...n-1] is defined to be an index such that A[i]
// = i. Given a sorted array of integers, write a method to find a magic index, if
// one exists, in array A. If not, return -1. If there are more than one magic inde
// x, return the smallest one.
//
// Example1:
//
//
// Input: nums = [0, 2, 3, 4, 5]
// Output: 0
//
//
// Example2:
//
//
// Input: nums = [1, 1, 1]
// Output: 1
//
//
// Note:
//
//
// 1 <= nums.length <= 1000000
//
// Related Topics 数组 二分查找
// 👍 33 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func findMagicIndex(nums []int) int {
	for i, val := range nums {
		if i == val {
			return i
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-07-31 08:56:59
