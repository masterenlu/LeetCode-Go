package Problem

import (
	"container/heap"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
type wrd struct {
	word string
	cnt  int
}
type hp []wrd

func (h hp) Len() int { return len(h) }
func (h hp) Less(a, b int) bool {
	s1, s2 := h[a], h[b]
	if s1.cnt == s2.cnt {
		return strings.Compare(s1.word, s2.word) == 1
	}
	return s1.cnt < s2.cnt
}
func (h hp) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
func (h *hp) Pop() interface{} {
	tmp := *h
	val := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return val
}
func (h *hp) Push(val interface{}) {
	*h = append(*h, val.(wrd))
}
func topKFrequent(words []string, k int) []string {
	dic := map[string]int{}
	for _, v := range words {
		dic[v]++
	}
	h := &hp{}
	for key, val := range dic {
		heap.Push(h, wrd{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(wrd).word
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

// 2021-05-20 09:45:05

// Given a non-empty list of words, return the k most frequent elements.
// Your answer should be sorted by frequency from highest to lowest. If two word
// s have the same frequency, then the word with the lower alphabetical order comes
// first.
//
// Example 1:
//
// Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
// Output: ["i", "love"]
// Explanation: "i" and "love" are the two most frequent words.
//    Note that "i" comes before "love" due to a lower alphabetical order.
//
//
//
// Example 2:
//
// Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"]
// , k = 4
// Output: ["the", "is", "sunny", "day"]
// Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
//    with the number of occurrence being 4, 3, 2 and 1 respectively.
//
//
//
// Note:
//
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// Input words contain only lowercase letters.
//
//
//
// Follow up:
//
// Try to solve it in O(n log k) time and O(n) extra space.
//
//
// 👍 264 👎 0
