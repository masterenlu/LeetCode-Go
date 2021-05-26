package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func reverseParentheses(s string) string {
	str := make([]byte, 0)
	stack := make([]int, 0)
	idx := 0
	for _, v := range s {
		if v == '(' {
			stack = append(stack, idx)
		} else if v == ')' {
			idxF := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			str = append(str[:idxF], arrReverse(str[idxF:])...)
		} else {
			str = append(str, byte(v))
			idx++
		}
	}
	return string(str)
}

func arrReverse(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i <= j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-26 08:24:35

// You are given a string s that consists of lower case English letters and brack
// ets.
//
// Reverse the strings in each pair of matching parentheses, starting from the i
// nnermost one.
//
// Your result should not contain any brackets.
//
//
// Example 1:
//
//
// Input: s = "(abcd)"
// Output: "dcba"
//
//
// Example 2:
//
//
// Input: s = "(u(love)i)"
// Output: "iloveu"
// Explanation: The substring "love" is reversed first, then the whole string is
// reversed.
//
//
// Example 3:
//
//
// Input: s = "(ed(et(oc))el)"
// Output: "leetcode"
// Explanation: First, we reverse the substring "oc", then "etco", and finally, t
// he whole string.
//
//
// Example 4:
//
//
// Input: s = "a(bcdefghijkl(mno)p)q"
// Output: "apmnolkjihgfedcbq"
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 2000
// s only contains lower case English characters and parentheses.
// It's guaranteed that all parentheses are balanced.
//
//
// 👍 87 👎 0
