package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func xorOperation(n int, start int) int {
	ans := start
	for i := 1; i < n; i++ {
		start += 2
		ans ^= start
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-07 09:46:37

// Given an integer n and an integer start.
//
// Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.le
// ngth.
//
// Return the bitwise XOR of all elements of nums.
//
//
// Example 1:
//
//
// Input: n = 5, start = 0
// Output: 8
// Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8)
// = 8.
// Where "^" corresponds to bitwise XOR operator.
//
//
// Example 2:
//
//
// Input: n = 4, start = 3
// Output: 8
// Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.
//
// Example 3:
//
//
// Input: n = 1, start = 7
// Output: 7
//
//
// Example 4:
//
//
// Input: n = 10, start = 5
// Output: 2
//
//
//
// Constraints:
//
//
// 1 <= n <= 1000
// 0 <= start <= 1000
// n == nums.length
//
// 👍 53 👎 0
