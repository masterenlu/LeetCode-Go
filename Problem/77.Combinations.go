package Problem

/*
	Go append 的参数并不是值传递，而是引用
*/

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	for i := 1; i <= n-k+1; i++ {
		temp := []int{i}
		backTrace(i+1, n, k, temp, &result)
	}
	return result
}

func backTrace(start, n, k int, path []int, result *[][]int) {
	if len(path) == k {
		// 这里需要注意，要深拷贝一个路径，直接 append path 在后面修改 path 时 result 也会跟着修改
		realPath := make([]int, k)
		copy(realPath, path)
		*result = append(*result, realPath)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		backTrace(i+1, n, k, path, result)
		path = path[:len(path)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-09-08 08:11:33

// Given two integers n and k, return all possible combinations of k numbers out
// of 1 ... n.
//
// You may return the answer in any order.
//
//
// Example 1:
//
//
// Input: n = 4, k = 2
// Output:
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
//
//
// Example 2:
//
//
// Input: n = 1, k = 1
// Output: [[1]]
//
//
//
// Constraints:
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// 👍 342 👎 0
