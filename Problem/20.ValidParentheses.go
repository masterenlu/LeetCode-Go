package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func isValid(s string) bool {
	stack := make([]byte, 0)

	for _, val := range s {
		l := len(stack) - 1
		switch {
		case val == '(', val == '[', val == '{':
			stack = append(stack, byte(val))
		case l == -1 && (val == ')' || val == ']' || val == '}'):
			return false
		case val == ')' && stack[l] == '(', val == ']' && stack[l] == '[', val == '}' && stack[l] == '{':
			stack = stack[:l]
		default:
			return false
		}
	}

	return len(stack) == 0
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-14 10:50:15

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
//
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
//
//
// Note that an empty string is also considered valid.
//
// Example 1:
//
//
// Input: "()"
// Output: true
//
//
// Example 2:
//
//
// Input: "()[]{}"
// Output: true
//
//
// Example 3:
//
//
// Input: "(]"
// Output: false
//
//
// Example 4:
//
//
// Input: "([)]"
// Output: false
//
//
// Example 5:
//
//
// Input: "{[]}"
// Output: true
//
//
// 👍 1768 👎 0
