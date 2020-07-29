package Problem

import (
	"strings"
)

//Given a string, you need to reverse the order of characters in each word withi
//n a sentence while still preserving whitespace and initial word order.
//
// Example 1:
//
//Input: "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
//
//
//
// Note:
//In the string, each word is separated by single space and there will not be an
//y extra space in the string.
// Related Topics 字符串
// 👍 205 👎 0

/*
	func main() {
		reverseWords("Let's go home.")
}*/

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	ans := make([]byte, 0)
	for _, subs := range ss {
		for i := len(subs) - 1; i >= 0; i-- {
			ans = append(ans, subs[i])
		}
		ans = append(ans, ' ')
	}
	ans = ans[:len(ans)-1]
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-28 13:16:54
