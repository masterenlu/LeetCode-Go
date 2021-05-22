package Problem

// leetcode submit region begin(Prohibit modification and deletion)
func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	ans := 0
	for _, val := range nums {
		ans ^= val
	}
	return ans == 0
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-22 08:28:28

// We are given non-negative integers nums[i] which are written on a chalkboard.
// Alice and Bob take turns erasing exactly one number from the chalkboard, with Al
// ice starting first. If erasing a number causes the bitwise XOR of all the elemen
// ts of the chalkboard to become 0, then that player loses. (Also, we'll say the b
// itwise XOR of one element is that element itself, and the bitwise XOR of no elem
// ents is 0.)
//
// Also, if any player starts their turn with the bitwise XOR of all the element
// s of the chalkboard equal to 0, then that player wins.
//
// Return True if and only if Alice wins the game, assuming both players play op
// timally.
//
//
// Example:
// Input: nums = [1, 1, 2]
// Output: false
// Explanation:
// Alice has two choices: erase 1 or erase 2.
// If she erases 1, the nums array becomes [1, 2]. The bitwise XOR of all the ele
// ments of the chalkboard is 1 XOR 2 = 3. Now Bob can remove any element he wants,
// because Alice will be the one to erase the last element and she will lose.
// If Alice erases 2 first, now nums becomes [1, 1]. The bitwise XOR of all the e
// lements of the chalkboard is 1 XOR 1 = 0. Alice will lose.
//
//
//
// Notes:
//
//
// 1 <= N <= 1000.
// 0 <= nums[i] <= 2^16.
//
//
//
//
// 👍 44 👎 0
