package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfFour(n int) bool {
	return n > 0 && n%3 == 1 && n&(n-1) == 0
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-31 09:17:56

// Given an integer n, return true if it is a power of four. Otherwise, return fa
// lse.
//
// An integer n is a power of four, if there exists an integer x such that n ==
// 4x.
//
//
// Example 1:
// Input: n = 16
// Output: true
// Example 2:
// Input: n = 5
// Output: false
// Example 3:
// Input: n = 1
// Output: true
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
// 👍 199 👎 0
