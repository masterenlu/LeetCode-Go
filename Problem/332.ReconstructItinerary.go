package Problem

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func findItinerary(tickets [][]string) []string {
	hash, ans, l := map[string][]string{}, make([]string, 0), len(tickets)

	// 构建出结点和入结点的 hash 表
	for i := 0; i < l; i++ {
		hash[tickets[i][0]] = append(hash[tickets[i][0]], tickets[i][1])
	}

	// 对每个入结点的数组值排序
	for _, val := range hash {
		sort.Strings(val)
	}

	// 递归搜索
	var dfs func(target string)
	dfs = func(target string) {
		for {
			// 当当前结点无出度时退出
			if val, ok := hash[target]; !ok || len(val) == 0 {
				break
			}
			temp := hash[target][0]
			hash[target] = hash[target][1:]
			dfs(temp)
		}
		ans = append(ans, target)
	}

	dfs("JFK")

	// 翻转
	for i := 0; i < (l+1)>>1; i++ {
		ans[i], ans[l-i] = ans[l-i], ans[i]
	}

	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-27 08:27:17

// Given a list of airline tickets represented by pairs of departure and arrival
// airports [from, to], reconstruct the itinerary in order. All of the tickets belo
// ng to a man who departs from JFK. Thus, the itinerary must begin with JFK.
//
// Note:
//
//
// If there are multiple valid itineraries, you should return the itinerary that
// has the smallest lexical order when read as a single string. For example, the i
// tinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
// All airports are represented by three capital letters (IATA code).
// You may assume all tickets form at least one valid itinerary.
// One must use all the tickets once and only once.
//
//
// Example 1:
//
//
// Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
// Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
//
//
// Example 2:
//
//
// Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
//
// Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
// Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL"
// ,"SFO"].
//              But it is larger in lexical order.
//
//
// 👍 169 👎 0
