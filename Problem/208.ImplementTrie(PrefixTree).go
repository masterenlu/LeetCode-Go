package Problem

// leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	isWord bool
	next   [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := Trie{isWord: false, next: [26]*Trie{nil}}
	return root
}

/** Inserts a word into the trie. */
func (root *Trie) Insert(word string) {
	pointer := root
	for idx := range word {
		if pointer.next[word[idx]-'a'] == nil {
			pointer.next[word[idx]-'a'] = &Trie{isWord: false, next: [26]*Trie{nil}}
		}
		pointer = pointer.next[word[idx]-'a']
	}
	pointer.isWord = true
}

// get node of given prefix
func (root *Trie) getPrefixNode(prefix string) *Trie {
	pointer := root
	for _, val := range prefix {
		if pointer.next[val-'a'] == nil {
			return nil
		}
		pointer = pointer.next[val-'a']
	}
	return pointer
}

/** Returns if the word is in the trie. */
func (root *Trie) Search(word string) bool {
	pointer := root.getPrefixNode(word)
	if pointer != nil && pointer.isWord {
		return true
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (root *Trie) StartsWith(prefix string) bool {
	pointer := root.getPrefixNode(prefix)
	if pointer != nil {
		return true
	} else {
		return false
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// leetcode submit region end(Prohibit modification and deletion)

// 2021-04-14 08:45:16

// A trie (pronounced as "try") or prefix tree is a tree data structure used to e
// fficiently store and retrieve keys in a dataset of strings. There are various ap
// plications of this data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
//
// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (i
// .e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously inser
// ted string word that has the prefix prefix, and false otherwise.
//
//
//
// Example 1:
//
//
// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]
//
// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True
//
//
//
// Constraints:
//
//
// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 104 calls in total will be made to insert, search, and startsWith
// .
//
//
// 👍 608 👎 0
