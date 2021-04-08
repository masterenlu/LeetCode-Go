package Problem

// You are a professional robber planning to rob houses along a street. Each hous
// e has a certain amount of money stashed, the only constraint stopping you from r
// obbing each of them is that adjacent houses have security system connected and i
// t will automatically contact the police if two adjacent houses were broken into
// on the same night.
//
// Given a list of non-negative integers representing the amount of money of eac
// h house, determine the maximum amount of money you can rob tonight without alert
// ing the police.
//
//
// Example 1:
//
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//              Total amount you can rob = 1 + 3 = 4.
//
//
// Example 2:
//
//
// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5
// (money = 1).
//              Total amount you can rob = 2 + 9 + 1 = 12.
//
//
//
// Constraints:
//
//
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
// Related Topics 动态规划
// 👍 982 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	for i, val := range nums {
		switch i {
		case 0:
			dp[i] = val
			ans = getMax(ans, dp[i])
		case 1:
			dp[i] = val
			ans = getMax(ans, dp[i])
		default:
			temp := 0
			for index := 0; index < i-1; index++ {
				temp = getMax(temp, dp[index])
			}
			dp[i] = temp + val
			ans = getMax(ans, dp[i])
		}
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

// 2020-08-05 09:42:26
