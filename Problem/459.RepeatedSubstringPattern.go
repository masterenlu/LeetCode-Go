package Problem

import (
	"fmt"
	"strings"
)

//goland:noinspection GoUnusedFunction
func main() {
	fmt.Println(repeatedSubstringPattern("ab"))
}

/*
	如果一个字符串是重复子串构成的，那么这个子串 s 的形式必然为 s'...s'
	将两个字符串拼接构成新字符串 ss，如果 s 可由子字符串重复，那么 ss 中必有一个子串 s''，使得 s'' 的索引在 0, len(s) 之间，ss的构成为 s'...s's'...s'
*/

// leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	temp := s + s
	l := len(s)
	if l < 2 {
		return false
	}
	s = strings.Join([]string{string(s[l-1]), s, string(s[0])}, "")
	return strings.Index(temp, s) != -1
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-24 09:06:35

// Given a non-empty string check if it can be constructed by taking a substring
// of it and appending multiple copies of the substring together. You may assume th
// e given string consists of lowercase English letters only and its length will no
// t exceed 10000.
//
//
//
// Example 1:
//
//
// Input: "abab"
// Output: True
// Explanation: It's the substring "ab" twice.
//
//
// Example 2:
//
//
// Input: "aba"
// Output: False
//
//
// Example 3:
//
//
// Input: "abcabcabcabc"
// Output: True
// Explanation: It's the substring "abc" four times. (And the substring "abcabc"
// twice.)
//
//
// 👍 266 👎 0
