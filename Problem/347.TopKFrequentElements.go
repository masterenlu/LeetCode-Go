package Problem

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}

	for _, val := range nums {
		count[val]++
	}

	ans := make([]int, 0)

	for j := 0; j < k; j++ {
		max, index := 0, 0
		for i, val := range count {
			if val >= max {
				index = i
				max = val
			}
		}
		ans = append(ans, index)
		delete(count, index)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-09-07 10:30:43

//Given a non-empty array of integers, return the k most frequent elements.
//
// Example 1:
//
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//
//
//
// Example 2:
//
//
//Input: nums = [1], k = 1
//Output: [1]
//
//
// Note:
//
//
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// Your algorithm's time complexity must be better than O(n log n), where n is t
//he array's size.
// It's guaranteed that the answer is unique, in other words the set of the top
//k frequent elements is unique.
// You can return the answer in any order.
//
//
// 👍 461 👎 0
