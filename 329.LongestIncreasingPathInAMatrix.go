package main

import "fmt"

//Given an integer matrix, find the length of the longest increasing path.
//
// From each cell, you can either move to four directions: left, right, up or do
//wn. You may NOT move diagonally or move outside of the boundary (i.e. wrap-aroun
//d is not allowed).
//
// Example 1:
//
//
//Input: nums =
//[
//  [9,9,4],
//  [6,6,8],
//  [2,1,1]
//]
//Output: 4
//Explanation: The longest increasing path is [1, 2, 6, 9].
//
//
// Example 2:
//
//
//Input: nums =
//[
//  [3,4,5],
//  [3,2,6],
//  [2,2,1]
//]
//Output: 4
//Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is
// not allowed.
//
// Related Topics 深度优先搜索 拓扑排序 记忆化
// 👍 222 👎 0

/*
	主要采用深度优先搜索的方法进行求解，需要注意的是如果暴力求解的话会在倒数第二个样例卡时间，所以要记录一下每个元素的最大路径
*/

// 测试

func main() {
	mat := [][]int{{1, 2}}
	fmt.Print(longestIncreasingPath(mat))
}

//leetcode submit region begin(Prohibit modification and deletion)

// 全局变量，矩阵的 row 和 column
var m, n int

// 深度优先搜索
func dfs(r, c int, matrix, val [][]int) int {
	// 对元素最大路径进行判断，如果不为 0，则值为最大路径值
	if val[r][c] != 0 {
		return val[r][c]
	}

	// 如果为 0，未初始化，将其置 1
	val[r][c]++

	// 按照四个方向进行探索
	if r+1 < m && matrix[r][c] < matrix[r+1][c] {
		val[r][c] = max(1+dfs(r+1, c, matrix, val), val[r][c])
	}
	if r-1 >= 0 && matrix[r][c] < matrix[r-1][c] {
		val[r][c] = max(1+dfs(r-1, c, matrix, val), val[r][c])
	}
	if c+1 < n && matrix[r][c] < matrix[r][c+1] {
		val[r][c] = max(1+dfs(r, c+1, matrix, val), val[r][c])
	}
	if c-1 >= 0 && matrix[r][c] < matrix[r][c-1] {
		val[r][c] = max(1+dfs(r, c-1, matrix, val), val[r][c])
	}

	return val[r][c]
}

func longestIncreasingPath(matrix [][]int) int {
	// 判断矩阵是否存在
	if len(matrix) == 0 || len(matrix) == 0 {
		return 0
	}

	// 获取 row 和 column
	m, n = len(matrix), len(matrix[0])

	// 定义最大路径矩阵
	val := make([][]int, m)
	for i := 0; i < m; i++ {
		val[i] = make([]int, n)
	}

	// 返回值
	maxL := 1

	// 对矩阵中的每个元素遍历判断
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxL = max(maxL, dfs(i, j, matrix, val))
		}
	}

	return maxL
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-26 08:41:48
