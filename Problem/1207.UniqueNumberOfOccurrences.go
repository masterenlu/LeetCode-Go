package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func uniqueOccurrences(arr []int) bool {
	hashCount := map[int]int{}
	hash := map[int]bool{}
	for _, v := range arr {
		hashCount[v]++
	}
	for _, v := range hashCount {
		if hash[v] {
			return false
		}
		hash[v] = true
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-10-28 10:41:57

// Given an array of integers arr, write a function that returns true if and only
// if the number of occurrences of each value in the array is unique.
//
//
// Example 1:
//
//
// Input: arr = [1,2,2,1,1,3]
// Output: true
// Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values
// have the same number of occurrences.
//
// Example 2:
//
//
// Input: arr = [1,2]
// Output: false
//
//
// Example 3:
//
//
// Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
// Output: true
//
//
//
// Constraints:
//
//
// 1 <= arr.length <= 1000
// -1000 <= arr[i] <= 1000
//
//
// 👍 66 👎 0
