package 剑指_Offer

//English description is not available for the problem. Please switch to Chinese
//. Related Topics 二分查找
// 👍 126 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {
	ans := numbers[0]

	for _, val := range numbers {
		if ans > val {
			return val
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-29 09:02:10
