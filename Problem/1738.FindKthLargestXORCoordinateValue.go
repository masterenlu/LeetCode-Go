package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{10, 9, 5}, {2, 0, 4}, {1, 0, 9}, {3, 4, 8}}

	kthLargestValue(arr, 10)
}

// leetcode submit region begin(Prohibit modification and deletion)
func kthLargestValue(matrix [][]int, k int) int {
	n, m := len(matrix)+1, len(matrix[0])+1
	pre := make([]int, n*m)
	for i := range matrix {
		for j, v := range matrix[i] {
			pre[i*m+j+m+1] = v ^ pre[i*m+j+m] ^ pre[i*m+j+1] ^ pre[i*m+j]
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pre)))
	fmt.Println(pre)
	return pre[k-1]
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-19 09:47:39

// You are given a 2D matrix of size m x n, consisting of non-negative integers.
// You are also given an integer k.
//
// The value of coordinate (a, b) of the matrix is the XOR of all matrix[i][j] w
// here 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).
//
// Find the kth largest value (1-indexed) of all the coordinates of matrix.
//
//
// Example 1:
//
//
// Input: matrix = [[5,2],[1,6]], k = 1
// Output: 7
// Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the larges
// t value.
//
// Example 2:
//
//
// Input: matrix = [[5,2],[1,6]], k = 2
// Output: 5
// Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest
// value.
//
// Example 3:
//
//
// Input: matrix = [[5,2],[1,6]], k = 3
// Output: 4
// Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd la
// rgest value.
//
// Example 4:
//
//
// Input: matrix = [[5,2],[1,6]], k = 4
// Output: 0
// Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which i
// s the 4th largest value.
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 1000
// 0 <= matrix[i][j] <= 106
// 1 <= k <= m * n
//
//
// 👍 31 👎 0
