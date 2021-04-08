package Problem

// Given a list of unique words, find all pairs of distinct indices (i, j) in the
// given list, so that the concatenation of the two words, i.e. words[i] + words[j
// ] is a palindrome.
//
// Example 1:
//
//
//
// Input: ["abcd","dcba","lls","s","sssll"]
// Output: [[0,1],[1,0],[3,2],[2,4]]
// Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]
//
//
//
// Example 2:
//
//
// Input: ["bat","tab","cat"]
// Output: [[0,1],[1,0]]
// Explanation: The palindromes are ["battab","tabbat"]
//
//
//
// Related Topics 字典树 哈希表 字符串
// 👍 85 👎 0

/*
	对于每一个字符串，建立关于其下标的哈希表
	对于每一个字符串，按照左右两个方向进行判定，对于每个方向依次递增字符串的子串长度，可以将字符串分为 s1, s2 两个部分
	判定 s1 的反转字符串是否存在，s2 是否为回文字符串，如果同时满足，则增加答案 { 字符串下标, 查找到的字符串下表 }
	判定 s2 的反转字符串和 s1 是否回文，同上
	注意需要判断空字符串的情况
*/

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func palindromePairs(words []string) [][]int {
	record := make(map[string]int)
	ans := make([][]int, 0)

	// 建立字符串和其下标的哈希表
	for index, val := range words {
		record[val] = index
	}

	// 对每一个字符串判断
	for indexFirst, valFirst := range words {
		for p := 0; p < len(valFirst); p++ {
			// 判断左侧子串的反转是否存在，右侧子串是否回文
			if indexSecond, flag := record[reverse(valFirst[:p])]; flag && isPalindrome(valFirst[p:]) && indexSecond != indexFirst {
				ans = append(ans, []int{indexFirst, indexSecond})
				// 子串为空情况
				if len(words[indexSecond]) == 0 {
					ans = append(ans, []int{indexSecond, indexFirst})
				}
			}
			// 判断右侧子串反转是否存在，左侧子串是否回文
			if indexSecond, flag := record[reverse(valFirst[p:])]; flag && isPalindrome(valFirst[:p]) && indexSecond != indexFirst {
				ans = append(ans, []int{indexSecond, indexFirst})
				// 子串为空情况
				if len(words[indexSecond]) == 0 {
					ans = append(ans, []int{indexFirst, indexSecond})
				}
			}
		}
	}
	return ans
}

// 判断字符串回文
func isPalindrome(s string) bool {
	length := len(s) - 1
	for i := 0; i < (length+1)>>1; i++ {
		if s[i] != s[length-i] {
			return false
		}
	}
	return true
}

// 获取字符串反转
func reverse(s string) string {
	revStr := []byte(s)
	length := len(s) - 1
	for i := 0; i < (length+1)>>1; i++ {
		revStr[i], revStr[length-i] = revStr[length-i], revStr[i]
	}
	return string(revStr)
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-06 08:56:56
