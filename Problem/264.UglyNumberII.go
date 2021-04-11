package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func nthUglyNumber(n int) int {
	u := []int{0, 0, 0}     // three pointers
	uglyN := make([]int, n) // ugly numbers
	uI, tempUgly := 0, 1    // pointer index and temp ugly number

	for idx := 0; idx < len(uglyN); idx++ {
		if idx == 0 {
			uglyN[0] = 1
		} else {
			tempUgly, uI = getMin(uglyN[u[0]]*2, uglyN[u[1]]*3, uglyN[u[2]]*5)
			if tempUgly != uglyN[idx-1] {
				uglyN[idx] = tempUgly
			} else {
				idx--
			}
			u[uI]++
		}
	}
	return uglyN[n-1]
}

func getMin(n1, n2, n3 int) (int, int) {
	if n1 < n2 && n1 < n3 {
		return n1, 0
	} else if n2 < n3 {
		return n2, 1
	} else {
		return n3, 2
	}
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-11 11:09:01

// Given an integer n, return the nth ugly number.
//
// Ugly number is a positive number whose prime factors only include 2, 3, and/o
// r 5.
//
//
// Example 1:
//
//
// Input: n = 10
// Output: 12
// Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10
// ugly numbers.
//
//
// Example 2:
//
//
// Input: n = 1
// Output: 1
// Explanation: 1 is typically treated as an ugly number.
//
//
//
// Constraints:
//
//
// 1 <= n <= 1690
//
//
// 👍 548 👎 0
