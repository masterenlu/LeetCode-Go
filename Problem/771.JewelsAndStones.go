package Problem

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func numJewelsInStones(J string, S string) int {
	hash := map[byte]bool{}
	for _, val := range J {
		hash[byte(val)] = true
	}

	ans := 0
	for _, val := range S {
		if hash[byte(val)] == true {
			ans += 1
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-10-02 22:51:39

// You're given strings J representing the types of stones that are jewels, and S
// representing the stones you have. Each character in S is a type of stone you ha
// ve. You want to know how many of the stones you have are also jewels.
//
// The letters in J are guaranteed distinct, and all characters in J and S are l
// etters. Letters are case sensitive, so "a" is considered a different type of sto
// ne from "A".
//
// Example 1:
//
//
// Input: J = "aA", S = "aAAbbbb"
// Output: 3
//
//
// Example 2:
//
//
// Input: J = "z", S = "ZZ"
// Output: 0
//
//
// Note:
//
//
// S and J will consist of letters and have length at most 50.
// The characters in J are distinct.
//
//
// 👍 587 👎 0
