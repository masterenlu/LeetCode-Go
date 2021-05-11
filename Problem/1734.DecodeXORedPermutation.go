package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func decode(encoded []int) []int {
	total := 0
	for i := 0; i < len(encoded)+1; i++ {
		total ^= i + 1
	}
	odd := 0
	for i := 1; i < len(encoded); i += 2 {
		odd ^= encoded[i]
	}
	ans := make([]int, len(encoded)+1)
	ans[0] = odd ^ total
	for i := 1; i < len(ans); i++ {
		ans[i] = encoded[i-1] ^ ans[i-1]
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-11 10:22:24

// There is an integer array perm that is a permutation of the first n positive i
// ntegers, where n is always odd.
//
// It was encoded into another integer array encoded of length n - 1, such that
// encoded[i] = perm[i] XOR perm[i + 1]. For example, if perm = [1,3,2], then encod
// ed = [2,1].
//
// Given the encoded array, return the original array perm. It is guaranteed tha
// t the answer exists and is unique.
//
//
// Example 1:
//
//
// Input: encoded = [3,1]
// Output: [1,2,3]
// Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]
//
//
// Example 2:
//
//
// Input: encoded = [6,5,4,6]
// Output: [2,4,1,5,3]
//
//
//
// Constraints:
//
//
// 3 <= n < 105
// n is odd.
// encoded.length == n - 1
//
//
// 👍 54 👎 0
