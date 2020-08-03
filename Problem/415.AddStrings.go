package Problem

//Given two non-negative integers num1 and num2 represented as string, return th
//e sum of num1 and num2.
//
// Note:
//
// The length of both num1 and num2 is < 5100.
// Both num1 and num2 contains only digits 0-9.
// Both num1 and num2 does not contain any leading zero.
// You must not use any built-in BigInteger library or convert the inputs to int
//eger directly.
//
// Related Topics 字符串
// 👍 194 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	a, b, c, d, flag1, flag2 := 0, 0, 0, 0, false, false
	ans := ""

	for i := 0; ; i++ {
		if i > len(num1)-1 {
			a = 0
			flag1 = true
		} else {
			a = int(num1[len(num1)-1-i] - '0')
		}

		if i > len(num2)-1 {
			b = 0
			flag2 = true
		} else {
			b = int(num2[len(num2)-1-i] - '0')
		}
		d = (a + b + c) % 10
		if flag1 && flag2 {
			break
		}
		ans = string(d+'0') + ans[0:]
		c = (a + b + c) / 10
	}
	if d != 0 {
		ans = string(d+'0') + ans[0:]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-03 08:31:16
