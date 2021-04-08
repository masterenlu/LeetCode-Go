package Problem

//goland:noinspection GoUnusedFunction
func main() {
	// s := "(name)is(age)yearsold"
	// knowledge := [][]string{{"name", "bob"}, {"age", "two"}}
	// s := "hi(name)"
	// knowledge := [][]string{{"a", "b"}}
	// s := "(a)(a)(a)aaa"
	// knowledge := [][]string{{"a", "yes"}}
	s := "(a)(b)"
	knowledge := [][]string{{"a", "b"}, {"b", "a"}}
	evaluate(s, knowledge)
}

// leetcode submit region begin(Prohibit modification and deletion)
func evaluate(s string, knowledge [][]string) string {
	hashmap := map[string]string{}
	for _, val := range knowledge {
		hashmap[val[0]] = val[1]
	}
	ans := ""
	for index := 0; index < len(s); index++ {
		if s[index] == '(' {
			secondIndex := findNextBracket(index, s)
			substr := s[index : secondIndex+1]
			key := substr[1 : len(substr)-1]
			value, ok := hashmap[key]
			if ok {
				ans += value
			} else {
				ans += "?"
			}
			index = index + len(substr) - 1
		} else {
			ans += s[index : index+1]
		}
	}
	return ans
}

func findNextBracket(firstIndex int, str string) int {
	for i := firstIndex + 1; i < len(str); i++ {
		if str[i] == ')' {
			return i
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-02 09:54:36

// You are given a string s that contains some bracket pairs, with each pair cont
// aining a non-empty key.
//
//
// For example, in the string "(name)is(age)yearsold", there are two bracket pai
// rs that contain the keys "name" and "age".
//
//
// You know the values of a wide range of keys. This is represented by a 2D stri
// ng array knowledge where each knowledge[i] = [keyi, valuei] indicates that key k
// eyi has a value of valuei.
//
// You are tasked to evaluate all of the bracket pairs. When you evaluate a brac
// ket pair that contains some key keyi, you will:
//
//
// Replace keyi and the bracket pair with the key's corresponding valuei.
// If you do not know the value of the key, you will replace keyi and the bracke
// t pair with a question mark "?" (without the quotation marks).
//
//
// Each key will appear at most once in your knowledge. There will not be any ne
// sted brackets in s.
//
// Return the resulting string after evaluating all of the bracket pairs.
//
//
// Example 1:
//
//
// Input: s = "(name)is(age)yearsold", knowledge = [["name","bob"],["age","two"]]
//
// Output: "bobistwoyearsold"
// Explanation:
// The key "name" has a value of "bob", so replace "(name)" with "bob".
// The key "age" has a value of "two", so replace "(age)" with "two".
//
//
// Example 2:
//
//
// Input: s = "hi(name)", knowledge = [["a","b"]]
// Output: "hi?"
// Explanation: As you do not know the value of the key "name", replace "(name)"
// with "?".
//
//
// Example 3:
//
//
// Input: s = "(a)(a)(a)aaa", knowledge = [["a","yes"]]
// Output: "yesyesyesaaa"
// Explanation: The same key can appear multiple times.
// The key "a" has a value of "yes", so replace all occurrences of "(a)" with "ye
// s".
// Notice that the "a"s not in a bracket pair are not evaluated.
//
//
// Example 4:
//
//
// Input: s = "(a)(b)", knowledge = [["a","b"],["b","a"]]
// Output: "ba"
//
//
// Constraints:
//
//
// 1 <= s.length <= 105
// 0 <= knowledge.length <= 105
// knowledge[i].length == 2
// 1 <= keyi.length, valuei.length <= 10
// s consists of lowercase English letters and round brackets '(' and ')'.
// Every open bracket '(' in s will have a corresponding close bracket ')'.
// The key in each bracket pair of s will be non-empty.
// There will not be any nested bracket pairs in s.
// keyi and valuei consist of lowercase English letters.
// Each keyi in knowledge is unique.
//
//
// 👍 4 👎 0
