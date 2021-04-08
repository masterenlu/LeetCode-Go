package Problem

// There are a total of numCourses courses you have to take, labeled from 0 to nu
// mCourses-1.
//
// Some courses may have prerequisites, for example to take course 0 you have to
// first take course 1, which is expressed as a pair: [0,1]
//
// Given the total number of courses and a list of prerequisite pairs, is it pos
// sible for you to finish all courses?
//
//
// Example 1:
//
//
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0. So it is poss
// ible.
//
//
// Example 2:
//
//
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
//              To take course 1 you should have finished course 0, and to take c
// ourse 0 you should
//              also have finished course 1. So it is impossible.
//
//
//
// Constraints:
//
//
// The input prerequisites is a graph represented by a list of edges, not adjace
// ncy matrices. Read more about how a graph is represented.
// You may assume that there are no duplicate edges in the input prerequisites.
//
// 1 <= numCourses <= 10^5
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序
// 👍 439 👎 0

/*
	拓扑排序，找出每个结点的入度和邻接点，当入度为 0 时，结点可选
	判断可选结点和所有结点的数量
*/

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 邻接矩阵
	adjaceMat := make([][]int, numCourses)
	// 入度
	inDegree := make([]int, numCourses)

	// 读取弧，构建邻接矩阵和入度表
	for _, val := range prerequisites {
		adjaceMat[val[1]] = append(adjaceMat[val[1]], val[0])
		inDegree[val[0]]++
	}

	// 可选队列
	var queue []int
	// 初始化
	for index, val := range inDegree {
		// 初始入度为 0 的结点
		if val == 0 {
			queue = append(queue, index)
		}
	}

	// 可选数量
	cnt := 0
	// 拓扑排序
	for len(queue) > 0 {
		tempIndex := queue[0]
		queue = queue[1:]
		cnt++
		// 更新矩阵
		for _, val := range adjaceMat[tempIndex] {
			inDegree[val]--
			if inDegree[val] == 0 {
				queue = append(queue, val)
			}
		}
	}

	return cnt == numCourses
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-04 09:17:56
