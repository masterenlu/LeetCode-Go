package main

func main() {
	countBinarySubstrings("110011111011011001111111111111001100101")
}

//leetcode submit region begin(Prohibit modification and deletion)
func countBinarySubstrings(s string) int {
	if len(s) < 2 {
		return 0
	}

	leftLen, ans := 1, 0
	for i := 1; i < len(s); i++ {
		if temp := int(s[i-1]); temp == int(s[i]) {
			leftLen++
		} else {
			for j := 0; j < leftLen && i+j < len(s); j++ {
				if int(s[i+j]) == int(s[i]) {
					ans++
				} else {
					break
				}
			}
			leftLen = 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-10 08:32:16

//Give a string s, count the number of non-empty (contiguous) substrings that ha
//ve the same number of 0's and 1's, and all the 0's and all the 1's in these subs
//trings are grouped consecutively.
//
// Substrings that occur multiple times are counted the number of times they occ
//ur.
//
// Example 1:
//
//Input: "00110011"
//Output: 6
//Explanation: There are 6 substrings that have equal number of consecutive 1's
//and 0's: "0011", "01", "1100", "10", "0011", and "01".
// Notice that some of these substrings repeat and are counted the number of tim
//es they occur.
// Also, "00110011" is not a valid substring because all the 0's (and 1's) are n
//ot grouped together.
//
//
//
// Example 2:
//
//Input: "10101"
//Output: 4
//Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal nu
//mber of consecutive 1's and 0's.
//
//
//
// Note:
// s.length will be between 1 and 50,000.
// s will only consist of "0" or "1" characters.
//
// 👍 179 👎 0
