package Problem

//Suppose an array sorted in ascending order is rotated at some pivot unknown to
// you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// You may assume no duplicate exists in the array.
//
// Example 1:
//
//
//Input: [3,4,5,1,2]
//Output: 1
//
//
// Example 2:
//
//
//Input: [4,5,6,7,0,1,2]
//Output: 0
//
// Related Topics 数组 二分查找
// 👍 224 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMin(nums []int) int {
	ans := nums[0]

	for _, v := range nums {
		if ans > v {
			return v
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-28 13:50:34
