package Problem

import (
	"strconv"
)

//goland:noinspection GoUnusedFunction
func main() {
	restoreIpAddresses("25525511135")
}

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	var ans []string

	for i := 1; i < 4 && i < len(s)-2; i++ {
		for j := i + 1; j < i+4 && j < len(s)-1; j++ {
			for k := j + 1; k < j+4 && k < len(s); k++ {
				if isLegal(s[:i]) && isLegal(s[i:j]) && isLegal(s[j:k]) && isLegal(s[k:]) {
					ans = append(ans, s[:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:])
				}
			}
		}
	}
	return ans
}

func isLegal(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	temp, _ := strconv.Atoi(s)
	if temp > 255 {
		return false
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-09 08:38:20

// Given a string containing only digits, restore it by returning all possible va
// lid IP address combinations.
//
// A valid IP address consists of exactly four integers (each integer is between
// 0 and 255) separated by single points.
//
// Example:
//
//
// Input: "25525511135"
// Output: ["255.255.11.135", "255.255.111.35"]
//
//
// 👍 330 👎 0
