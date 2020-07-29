package 剑指_Offer

//English description is not available for the problem. Please switch to Chinese
//. Related Topics 动态规划
// 👍 50 👎 0
/*
	同 Problem.64
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxValue(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	valMat := make([][]int, n)

	for i := 0; i < n; i++ {
		valMat[i] = make([]int, m)
	}
	valMat[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		valMat[0][i] = grid[0][i] + valMat[0][i-1]
	}
	for i := 1; i < n; i++ {
		valMat[i][0] = grid[i][0] + valMat[i-1][0]
		for j := 1; j < m; j++ {
			valMat[i][j] = grid[i][j] + getMax(valMat[i-1][j], valMat[i][j-1])
		}
	}
	return valMat[n-1][m-1]
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-29 14:35:20
