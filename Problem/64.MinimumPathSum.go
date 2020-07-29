package Problem

//Given a m x n grid filled with non-negative numbers, find a path from top left
// to bottom right which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
// Example:
//
//
//Input:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.
//
// Related Topics 数组 动态规划
// 👍 544 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	val := make([][]int, len(grid))

	for i := 0; i < len(val); i++ {
		val[i] = make([]int, len(grid[0]))
	}

	val[0][0] = grid[0][0]

	for i := 1; i < len(val[0]); i++ {
		val[0][i] = val[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(val); i++ {
		val[i][0] = val[i-1][0] + grid[i][0]

		for j := 1; j < len(val[0]); j++ {
			val[i][j] = grid[i][j] + getMin(val[i-1][j], val[i][j-1])
		}
	}
	return val[len(grid)-1][len(grid[0])-1]

}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-23 09:36:39
