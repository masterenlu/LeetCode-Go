package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-30 09:44:21

// Given an integer n, return true if it is a power of two. Otherwise, return fal
// se.
//
// An integer n is a power of two, if there exists an integer x such that n == 2
// x.
//
//
// Example 1:
//
//
// Input: n = 1
// Output: true
// Explanation: 20 = 1
//
//
// Example 2:
//
//
// Input: n = 16
// Output: true
// Explanation: 24 = 16
//
//
// Example 3:
//
//
// Input: n = 3
// Output: false
//
//
// Example 4:
//
//
// Input: n = 4
// Output: true
//
//
// Example 5:
//
//
// Input: n = 5
// Output: false
//
//
//
// Constraints:
//
//
// -231 <= n <= 231 - 1
//
//
//
// Follow up: Could you solve it without loops/recursion?
// 👍 331 👎 0
