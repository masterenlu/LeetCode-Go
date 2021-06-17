package Problem

import (
	"fmt"
	"regexp"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isNumber(s string) bool {
	s = strings.ToLower(s)
	nums := strings.Split(s, "e")
	fmt.Println(len(nums))
	switch len(nums) {
	case 1:
		return isInt(nums[0]) || isDecimal(nums[0])
	case 2:
		return (isInt(nums[0]) || isDecimal(nums[0])) && isInt(nums[1])
	default:
		return false
	}
}

func isInt(s string) bool {
	flag, _ := regexp.MatchString("^[\\+\\-]?[0-9]+$", s)
	return flag
}
func isDecimal(s string) bool {
	nums := strings.Split(s, ".")
	switch len(nums) {
	case 2:
		flag1, _ := regexp.MatchString("^[\\+\\-]?[0-9]*$", nums[0])
		flag2, _ := regexp.MatchString("^[0-9]*$", nums[1])
		if flag1 {
			nums[0] = strings.Replace(nums[0], "+", "", -1)
			nums[0] = strings.Replace(nums[0], "-", "", -1)
		}
		return flag1 && flag2 && len(nums[1])+len(nums[0]) >= 1
	default:
		return false
	}
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-06-17 10:43:37

// A valid number can be split up into these components (in order):
//
//
// A decimal number or an integer.
// (Optional) An 'e' or 'E', followed by an integer.
//
//
// A decimal number can be split up into these components (in order):
//
//
// (Optional) A sign character (either '+' or '-').
// One of the following formats:
//
// One or more digits, followed by a dot '.'.
// One or more digits, followed by a dot '.', followed by one or more digits.
// A dot '.', followed by one or more digits.
//
//
//
//
// An integer can be split up into these components (in order):
//
//
// (Optional) A sign character (either '+' or '-').
// One or more digits.
//
//
// For example, all the following are valid numbers: ["2", "0089", "-0.1", "+3.1
// 4", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"], w
// hile the following are not valid numbers: ["abc", "1a", "1e", "e3", "99e2.5", "-
// -6", "-+3", "95a54e53"].
//
// Given a string s, return true if s is a valid number.
//
//
// Example 1:
//
//
// Input: s = "0"
// Output: true
//
//
// Example 2:
//
//
// Input: s = "e"
// Output: false
//
//
// Example 3:
//
//
// Input: s = "."
// Output: false
//
//
// Example 4:
//
//
// Input: s = ".1"
// Output: true
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 20
// s consists of only English letters (both uppercase and lowercase), digits (0-
// 9), plus '+', minus '-', or dot '.'.
//
//
// 👍 214 👎 0
