package main

// leetcode submit region begin(Prohibit modification and deletion)

// ==================== Brute-Force =====================
// func strStr(haystack string, needle string) int {
// 	if haystack == needle {
// 		return 0
// 	}
// 	l := len(needle)
// 	for idx := 0; idx <= len(haystack)-l; idx++ {
// 		subs := haystack[idx : idx+l]
// 		if subs == needle {
// 			return idx
// 		}
// 	}
// 	return -1
// }
// ===================== KMP ======================
func strStr(haystack string, needle string) int {
	// TODO
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-09 09:52:39

// Implement strStr().
//
// Return the index of the first occurrence of needle in haystack, or -1 if need
// le is not part of haystack.
//
// Clarification:
//
// What should we return when needle is an empty string? This is a great questio
// n to ask during an interview.
//
// For the purpose of this problem, we will return 0 when needle is an empty str
// ing. This is consistent to C's strstr() and Java's indexOf().
//
//
// Example 1:
// Input: haystack = "hello", needle = "ll"
// Output: 2
// Example 2:
// Input: haystack = "aaaaa", needle = "bba"
// Output: -1
// Example 3:
// Input: haystack = "", needle = ""
// Output: 0
//
//
// Constraints:
//
//
// 0 <= haystack.length, needle.length <= 5 * 104
// haystack and needle consist of only lower-case English characters.
//
//
// 👍 771 👎 0
