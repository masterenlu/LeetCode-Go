package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) bool {
	highIndex := len(nums) - 1
	for idx := 0; idx < len(nums)-1; idx++ {
		if nums[idx] > nums[idx+1] {
			highIndex = idx
			break
		}
	}
	var ans bool
	if highIndex == len(nums)-1 {
		ans = binSearch(0, highIndex, target, nums)
	} else {
		ans = binSearch(0, highIndex, target, nums) || binSearch(highIndex+1, len(nums)-1, target, nums)
	}
	return ans
}

func binSearch(low, high, target int, nums []int) bool {
	if nums[low] > target || nums[high] < target {
		return false
	}
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-07 09:03:20

// There is an integer array nums sorted in non-decreasing order (not necessarily
// with distinct values).
//
// Before being passed to your function, nums is rotated at an unknown pivot ind
// ex k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1]
// , ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0
// ,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,
// 2,4,4].
//
// Given the array nums after the rotation and an integer target, return true if
// target is in nums, or false if it is not in nums.
//
//
// Example 1:
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true
// Example 2:
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false
//
//
// Constraints:
//
//
// 1 <= nums.length <= 5000
// -104 <= nums[i] <= 104
// nums is guaranteed to be rotated at some pivot.
// -104 <= target <= 104
//
//
//
// Follow up: This problem is the same as Search in Rotated Sorted Array, where n
// ums may contain duplicates. Would this affect the runtime complexity? How and wh
// y?
// 👍 329 👎 0
