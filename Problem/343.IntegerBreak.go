package Problem

import "math"

// Given a positive integer n, break it into the sum of at least two positive int
// egers and maximize the product of those integers. Return the maximum product you
// can get.
//
// Example 1:
//
//
//
// Input: 2
// Output: 1
// Explanation: 2 = 1 + 1, 1 × 1 = 1.
//
//
// Example 2:
//
//
// Input: 10
// Output: 36
// Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
//
// Note: You may assume that n is not less than 2 and not larger than 58.
//
// Related Topics 数学 动态规划
// 👍 283 👎 0

/*
	当乘数最接近时乘积最大，只需要比较分成 2~n-1 份时获得乘积的大小即可
*/

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func integerBreak(n int) int {
	ans := 1

	for i := 2; i < n; i++ {
		// 分成 2~n-1 份
		temp := 1

		if n%i != 0 {
			// 不能平均分配时，将余数尽可能多的分配到乘数上
			temp = int(math.Pow(float64(n/i+1), float64(n%i))) * int(math.Pow(float64(n/i), float64(i-n%i)))
		} else {
			temp = int(math.Pow(float64(n/i), float64(i)))
		}
		ans = getMax(ans, temp)
	}
	return ans
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-07-30 09:34:00
