package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func hammingDistance(x int, y int) int {
	ans := 0
	// 每次循环消除最右侧1，循环次数为答案
	for z := x ^ y; z != 0; z &= z - 1 {
		ans++
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-27 08:53:42

// The Hamming distance between two integers is the number of positions at which
// the corresponding bits are different.
//
// Given two integers x and y, return the Hamming distance between them.
//
//
// Example 1:
//
//
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//       ↑   ↑
// The above arrows point to positions where the corresponding bits are different
// .
//
//
// Example 2:
//
//
// Input: x = 3, y = 1
// Output: 1
//
//
//
// Constraints:
//
//
// 0 <= x, y <= 231 - 1
//
//
// 👍 426 👎 0
