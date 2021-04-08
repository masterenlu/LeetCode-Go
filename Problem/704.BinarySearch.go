package Problem

// Given a sorted (in ascending order) integer array nums of n elements and a tar
// get value, write a function to search target in nums. If target exists, then ret
// urn its index, otherwise return -1.
//
//
// Example 1:
//
//
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
//
//
//
// Example 2:
//
//
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
//
//
//
//
// Note:
//
//
// You may assume that all elements in nums are unique.
// n will be in the range [1, 10000].
// The value of each element in nums will be in the range [-9999, 9999].
//
// Related Topics 二分查找
// 👍 141 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) >> 1
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-07-28 09:58:02
