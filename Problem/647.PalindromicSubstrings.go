package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func countSubstrings(s string) int {
	ans, l := 0, len(s)

	for i := 0; i < l; i++ {
		ans += getPalindromicCount(i, i, s)
		ans += getPalindromicCount(i, i+1, s)
	}
	return ans
}

func getPalindromicCount(l, r int, s string) int {
	count := 0
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			count++
			l--
			r++
		} else {
			break
		}
	}
	return count
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-19 09:15:06

// Given a string, your task is to count how many palindromic substrings in this
// string.
//
// The substrings with different start indexes or end indexes are counted as dif
// ferent substrings even they consist of same characters.
//
// Example 1:
//
//
// Input: "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
//
//
//
//
// Example 2:
//
//
// Input: "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
//
//
//
//
// Note:
//
//
// The input string length won't exceed 1000.
//
//
//
// 👍 316 👎 0
