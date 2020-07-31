package Problem

//All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T,
//for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify
// repeated sequences within the DNA.
//
// Write a function to find all the 10-letter-long sequences (substrings) that o
//ccur more than once in a DNA molecule.
//
// Example:
//
//
//Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//
//Output: ["AAAAACCCCC", "CCCCCAAAAA"]
//
// Related Topics 位运算 哈希表
// 👍 106 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatedDnaSequences(s string) []string {
	hashmap := map[string]int{}
	ans := make([]string, 0)

	for i := 0; i < len(s)-9; i++ {
		temp := s[i : i+10]
		hashmap[temp]++
		if hashmap[temp] == 2 {
			ans = append(ans, temp)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

//2020-07-31 09:38:17
