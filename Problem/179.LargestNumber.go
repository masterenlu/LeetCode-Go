package Problem

import (
	"bytes"
	"sort"
	"strconv"
)

// leetcode submit region begin(Prohibit modification and deletion)
type Int []int

func (a Int) Len() int {
	return len(a)
}

func (a Int) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Int) Less(i, j int) bool {
	str1 := strconv.Itoa(a[i])
	str2 := strconv.Itoa(a[j])
	n1, _ := strconv.Atoi(str1 + str2)
	n2, _ := strconv.Atoi(str2 + str1)
	return n1 > n2
}

func largestNumber(nums []int) string {
	sort.Sort(Int(nums))
	var ans bytes.Buffer
	for _, val := range nums {
		ans.WriteString(strconv.Itoa(val))
	}
	if ansn, _ := strconv.Atoi(ans.String()); ansn == 0 {
		return "0"
	}
	return ans.String()
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-12 09:48:13

// Given a list of non-negative integers nums, arrange them such that they form t
// he largest number.
//
// Note: The result may be very large, so you need to return a string instead of
// an integer.
//
//
// Example 1:
//
//
// Input: nums = [10,2]
// Output: "210"
//
//
// Example 2:
//
//
// Input: nums = [3,30,34,5,9]
// Output: "9534330"
//
//
// Example 3:
//
//
// Input: nums = [1]
// Output: "1"
//
//
// Example 4:
//
//
// Input: nums = [10]
// Output: "10"
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 109
//
//
// 👍 546 👎 0
