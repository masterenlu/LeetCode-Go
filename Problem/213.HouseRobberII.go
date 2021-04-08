package Problem

// You are a professional robber planning to rob houses along a street. Each hous
// e has a certain amount of money stashed. All houses at this place are arranged i
// n a circle. That means the first house is the neighbor of the last one. Meanwhil
// e, adjacent houses have security system connected and it will automatically cont
// act the police if two adjacent houses were broken into on the same night.
//
// Given a list of non-negative integers representing the amount of money of eac
// h house, determine the maximum amount of money you can rob tonight without alert
// ing the police.
//
// Example 1:
//
//
// Input: [2,3,2]
// Output: 3
// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money =
// 2),
//              because they are adjacent houses.
//
//
// Example 2:
//
//
// Input: [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//              Total amount you can rob = 1 + 3 = 4.
// Related Topics 动态规划
// 👍 336 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp1, dp2 := 0, 0
	for i, val := range nums {
		switch i {
		case 0:
			dp1 = val
			dp2 = val
		case 1:
			dp2 = getMax(val, dp2)
		case len(nums) - 1:
			break
		default:
			temp := dp2
			dp2 = getMax(dp2, dp1+val)
			dp1 = temp
		}
	}
	ans := dp2

	for i, val := range nums {
		switch i {
		case 0:
			continue
		case 1:
			dp1 = val
			dp2 = val
		case 2:
			dp2 = getMax(val, dp2)
		default:
			temp := dp2
			dp2 = getMax(dp2, dp1+val)
			dp1 = temp
		}
	}

	ans = getMax(ans, dp2)
	return ans
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-05 09:43:04
