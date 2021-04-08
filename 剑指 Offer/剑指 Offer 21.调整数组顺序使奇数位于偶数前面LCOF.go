package 剑指_Offer

// English description is not available for the problem. Please switch to Chinese
// .
// 👍 33 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]%2 == 0 && nums[r]%2 == 1 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		} else if nums[r]%2 == 0 {
			r--
		} else if nums[l]%2 == 1 {
			l++
		}
	}

	return nums
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-07-31 11:34:22
