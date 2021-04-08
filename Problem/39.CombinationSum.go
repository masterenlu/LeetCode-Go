package Problem

import (
	"sort"
)

// leetcode submit region begin(Prohibit modification and deletion)

//goland:noinspection GoUnusedFunction
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)
	var dfs func(index, target int)
	dfs = func(index, target int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}
		dfs(index+1, target)
		if target-candidates[index] >= 0 {
			path = append(path, candidates[index])
			dfs(index, target-candidates[index])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return result
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-09-09 07:23:26

// Given a set of candidate numbers (candidates) (without duplicates) and a targe
// t number (target), find all unique combinations in candidates where the candidat
// e numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of ti
// mes.
//
// Note:
//
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
//
//
// Example 1:
//
//
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//  [7],
//  [2,2,3]
// ]
//
//
// Example 2:
//
//
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]
//
//
//
// Constraints:
//
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// Each element of candidate is unique.
// 1 <= target <= 500
//
//
// 👍 853 👎 0
