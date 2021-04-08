package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func shortestPalindrome(s string) string {
	var ans []byte

	for i := len(s) - 1; i >= 0; i-- {
		if isPalindrome(s[:i+1]) {
			break
		} else {
			ans = append(ans, s[i])
			// ans += string(s[i])
		}
	}
	return string(ans) + s
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)>>1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-29 09:13:24

// Given a string s, you are allowed to convert it to a palindrome by adding char
// acters in front of it. Find and return the shortest palindrome you can find by p
// erforming this transformation.
//
// Example 1:
//
//
// Input: "aacecaaa"
// Output: "aaacecaaa"
//
//
// Example 2:
//
//
// Input: "abcd"
// Output: "dcbabcd"
// 👍 198 👎 0
