package Problem

// Suppose an array sorted in ascending order is rotated at some pivot unknown to
// you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// The array may contain duplicates.
//
// Example 1:
//
//
// Input: [1,3,5]
// Output: 1
//
// Example 2:
//
//
// Input: [2,2,2,0,1]
// Output: 0
//
// Note:
//
//
// This is a follow up problem to Find Minimum in Rotated Sorted Array.
// Would allow duplicates affect the run-time complexity? How and why?
//
// Related Topics 数组 二分查找
// 👍 146 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2

		if nums[m] < nums[r] {
			r = m
		} else if nums[m] == nums[r] {
			r--
		} else {
			l = m + 1
		}
	}

	return nums[l]
}

// leetcode submit region end(Prohibit modification and deletion)
