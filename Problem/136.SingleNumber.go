package Problem

// Given a non-empty array of integers, every element appears twice except for on
// e. Find that single one.
//
// Note:
//
// Your algorithm should have a linear runtime complexity. Could you implement i
// t without using extra memory?
//
// Example 1:
//
//
// Input: [2,2,1]
// Output: 1
//
//
// Example 2:
//
//
// Input: [4,1,2,1,2]
// Output: 4
//
// Related Topics 位运算 哈希表
// 👍 1383 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-07-24 15:53:29
