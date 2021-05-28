package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func totalHammingDistance(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < 30; i++ {
		cnt := 0
		for _, val := range nums {
			cnt += (val >> i) & 1
		}
		ans += cnt * (n - cnt)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-28 09:25:13

// The Hamming distance between two integers is the number of positions at which
// the corresponding bits are different.
//
// Given an integer array nums, return the sum of Hamming distances between all
// the pairs of the integers in nums.
//
//
// Example 1:
//
//
// Input: nums = [4,14,2]
// Output: 6
// Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 001
// 0 (just
// showing the four bits relevant in this case).
// The answer will be:
// HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 +
// 2 + 2 = 6.
//
//
// Example 2:
//
//
// Input: nums = [4,14,4]
// Output: 4
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 105
// 0 <= nums[i] <= 109
//
//
// 👍 166 👎 0
