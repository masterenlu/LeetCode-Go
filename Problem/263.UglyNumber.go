package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	prime := []int{2, 3, 5}
	flag := false
	for !flag && n != 1 {
		for idx := 0; idx < 3; idx++ {
			if n%prime[idx] == 0 {
				n /= prime[idx]
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
		flag = false
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-10 09:12:35

// Given an integer n, return true if n is an ugly number.
//
// Ugly number is a positive number whose prime factors only include 2, 3, and/o
// r 5.
//
//
// Example 1:
//
//
// Input: n = 6
// Output: true
// Explanation: 6 = 2 × 3
//
// Example 2:
//
//
// Input: n = 8
// Output: true
// Explanation: 8 = 2 × 2 × 2
//
//
// Example 3:
//
//
// Input: n = 14
// Output: false
// Explanation: 14 is not ugly since it includes another prime factor 7.
//
//
// Example 4:
//
//
// Input: n = 1
// Output: true
// Explanation: 1 is typically treated as an ugly number.
//
//
//
// Constraints:
//
//
// -231 <= n <= 231 - 1
//
//
// 👍 201 👎 0
