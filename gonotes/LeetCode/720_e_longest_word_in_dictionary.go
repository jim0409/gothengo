package main

import "sort"

/*
lexicographical: 詞典上的

Given a list of strings words representing an English Dictionary,
find `the longest word` in `words` that can be built one character at a time
by other words in words.

If there is more than one possible answer,
return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.


Example 1:
Input:
words = ["w","wo","wor","worl", "world"]
Output: "world"
Explanation:
The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".


Example 2:
Input:
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
Output: "apple"
Explanation:
Both "apply" and "apple" can be built from other words in the dictionary.
However, "apple" is lexicographically smaller than "apply".

Note:
All the strings in the input will only contain lowercase letters.
The length of words will be in the range [1, 1000].
The length of words[i] will be in the range [1, 30].
*/

// - https://leetcode.com/problems/longest-word-in-dictionary/discuss/578778/go-clean-code-4ms-beats-100
func longestWord(words []string) string {
	sort.Strings(words)
	var ret string
	set := make(map[string]struct{}, len(words))
	for _, word := range words {
		if _, ok := set[word[:len(word)-1]]; !ok && len(word) != 1 {
			continue
		}
		set[word] = struct{}{}
		if len(word) > len(ret) {
			ret = word
		}
	}
	return ret
}

func main() {

}

/*
solution:
// - https://leetcode.com/problems/longest-word-in-dictionary/discuss/546052/go-lang
func longestWord(words []string) string {
	obj := Constructor()
	for i := 0; i < len(words); i++ {
		obj.Insert(words[i])
	}

	maxword := ""
	flag := 0

	//use dfs to implement
	for i := 0; i < len(words); i++ {
		for j := 1; j < len(words[i]); j++ {
			param_2 := obj.Search(words[i][:j])
			if param_2 == false {
				flag = 1
				break
			}
		}
		if flag == 0 {
			if len(maxword) < len(words[i]) {
				maxword = words[i]
			} else if len(maxword) == len(words[i]) {
				if maxword[0] < words[i][0] {
				} else {
					for k := 0; k < len(maxword); k++ {
						if maxword[k] > words[i][k] {
							maxword = words[i]
							break
						}
					}
				}
			}
		}
		flag = 0
	}
	return maxword
}

type Trie struct {
	isKey        bool
	isVisited    bool //是否被访问过
	childrenNode [26]*Trie
}

// Initialize your data structure here.
func Constructor() Trie {
	var Tnode Trie
	Tnode.isKey = false
	Tnode.isVisited = false
	for i := 0; i < 26; i++ {
		Tnode.childrenNode[i] = nil
	}
	return Tnode
}

// Inserts a word into the trie.
func (this *Trie) Insert(word string) {

	for i := 0; i < len(word); i++ {
		if (*this).childrenNode[word[i]-'a'] == nil {
			temp := Constructor()
			(*this).childrenNode[word[i]-'a'] = &temp
		}
		this = (*this).childrenNode[word[i]-'a']
	}
	(*this).isKey = true

}

// Returns if the word is in the trie.
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if (*this).childrenNode[word[i]-'a'] == nil {
			return false
		}
		this = (*this).childrenNode[word[i]-'a']
	}
	if (*this).isKey == true {
		return true
	} else {
		return false
	}

}

// Returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if (*this).childrenNode[prefix[i]-'a'] == nil {
			return false
		}
		this = (*this).childrenNode[prefix[i]-'a']
	}
	return true

}
*/
