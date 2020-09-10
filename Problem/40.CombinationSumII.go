package Problem

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)
	var backTrace func(index, target int)
	backTrace = func(index, target int) {
		if target == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}

		for i := index; i < len(candidates); i++ {
			if i != index && candidates[i] == candidates[i-1] {
				continue
			} else if target-candidates[i] < 0 {
				return
			} else {
				path = append(path, candidates[i])
				backTrace(i+1, target-candidates[i])
				path = path[:len(path)-1]
			}
		}
	}
	backTrace(0, target)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-09-10 08:40:37

//Given a collection of candidate numbers (candidates) and a target number (targ
//et), find all unique combinations in candidates where the candidate numbers sums
// to target.
//
// Each number in candidates may only be used once in the combination.
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
//Input: candidates = [10,1,2,7,6,1,5], target = 8,
//A solution set is:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//
//
// Example 2:
//
//
//Input: candidates = [2,5,2,1,2], target = 5,
//A solution set is:
//[
//  [1,2,2],
//  [5]
//]
//
//
// 👍 358 👎 0
