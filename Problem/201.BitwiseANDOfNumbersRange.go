package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func rangeBitwiseAnd(m int, n int) int {
	bit := 0
	for m < n {
		m, n = m>>1, n>>1
		bit++
	}
	return m << bit
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-23 08:50:37

// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND o
// f all numbers in this range, inclusive.
//
// Example 1:
//
//
// Input: [5,7]
// Output: 4
//
//
// Example 2:
//
//
// Input: [0,1]
// Output: 0
// 👍 158 👎 0
