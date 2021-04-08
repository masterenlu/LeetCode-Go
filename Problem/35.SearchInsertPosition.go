package Problem

// Given a sorted array and a target value, return the index if the target is fou
// nd. If not, return the index where it would be if it were inserted in order.
//
// You may assume no duplicates in the array.
//
// Example 1:
//
//
// Input: [1,3,5,6], 5
// Output: 2
//
//
// Example 2:
//
//
// Input: [1,3,5,6], 2
// Output: 1
//
//
// Example 3:
//
//
// Input: [1,3,5,6], 7
// Output: 4
//
//
// Example 4:
//
//
// Input: [1,3,5,6], 0
// Output: 0
//
// Related Topics 数组 二分查找
// 👍 629 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	index := len(nums)
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
			index = m
		}
	}
	return index
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-05 11:40:15
