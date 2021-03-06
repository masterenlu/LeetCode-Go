package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for index, val := range nums {
		if v, ok := hash[target-val]; ok {
			return []int{v, index}
		}
		hash[val] = index
	}
	return []int{}
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-10-03 10:05:12

// Given an array of integers nums and an integer target, return indices of the t
// wo numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may n
// ot use the same element twice.
//
// You can return the answer in any order.
//
//
// Example 1:
//
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
//
//
// Example 2:
//
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
//
//
// Example 3:
//
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
//
//
//
// Constraints:
//
//
// 2 <= nums.length <= 105
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.
//
//
// 👍 9269 👎 0
