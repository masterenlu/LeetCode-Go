package main

//Given an array which consists of non-negative integers and an integer m, you c
//an split the array into m non-empty continuous subarrays. Write an algorithm to
//minimize the largest sum among these m subarrays.
//
//
// Note:
//If n is the length of array, assume the following constraints are satisfied:
//
// 1 ≤ n ≤ 1000
// 1 ≤ m ≤ min(50, n)
//
//
//
// Examples:
//
//Input:
//nums = [7,2,5,10,8]
//m = 2
//
//Output:
//18
//
//Explanation:
//There are four ways to split nums into two subarrays.
//The best way is to split it into [7,2,5] and [10,8],
//where the largest sum among the two subarrays is only 18.
//
// Related Topics 二分查找 动态规划
// 👍 201 👎 0

/*
	由题目可知，分块数组和最大值中的最小值一定在 ( max(nums), sum(nums) ) 中间，这里运用二分搜索对每一个在这个区间中的中间值进行检测，最终结果就是答案
	需要注意的是，这道题中的二分搜索不是指对下标进行二分搜索，而是对值进行二分搜索，并对取得的每一个中间值进行检验，直到收敛
*/

//leetcode submit region begin(Prohibit modification and deletion)
func splitArray(nums []int, m int) int {
	l, r := 0, 0

	// 左右区间
	for _, val := range nums {
		r += val
		if l < val {
			l = val
		}
	}

	// 值的二分搜索
	for l < r {
		mid := l + (r-l)/2 // 获取中间值
		sum, cnt := 0, 1   // sum 为一个分块的和，cnt 为分块数量

		// 计算两个数值
		for _, val := range nums {
			if sum+val > mid {
				// 如果当前分块加上要添加的元素的和大于要测试的中间值，当前分块结束，当前元素作为下一分块的初始值
				sum = val
				cnt++ // 分块数量增加
			} else {
				sum += val
			}
		}

		// 对获得的 cnt 进行检测
		if cnt > m {
			// 如果获得的 cnt 大于要切分的数量，意味着 mid 过小，收缩左边界
			l = mid + 1
		} else {
			// 否则 mid 过大，收缩右边界
			r = mid
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-25 09:25:25
