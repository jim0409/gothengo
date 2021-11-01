package main

import "strings"

/*
We are given two sentences s1 and s2.
(A sentence is a string of space separated words. Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.


Example 1:
Input: s1 = "this apple is sweet", s2 = "this apple is sour"
Output: ["sweet","sour"]

Example 2:
Input: s1 = "apple apple", s2 = "banana"
Output: ["banana"]


Note:
0 <= s1.length <= 200
0 <= s2.length <= 200
s1 and s2 both contain only spaces and lowercase letters.
*/

// 在一系列的文字中，找尋出現頻率最低的文字
func uncommonFromSentences(A string, B string) []string {
	// 將所有字放進一個array中
	allString := append(strings.Split(A, " "), strings.Split(B, " ")...)
	ans := allString[:0]
	var wordsFreqMap = make(map[string]int)

	// 計算所有字array中，出現的頻率字次數
	for _, v := range allString {
		if k, ok := wordsFreqMap[v]; ok {
			k++
			wordsFreqMap[v] = k
		} else {
			wordsFreqMap[v] = 1
		}
	}

	// 找尋頻率字中，出現次數為1的字，該字即為最低頻字
	for k, v := range wordsFreqMap {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/uncommon-words-from-two-sentences/discuss/581639/go-clean-code-0ms-beats-100

func uncommonFromSentences(A string, B string) []string {
	data := make(map[string]int)
	for _, word := range strings.Fields(A + " " + B) {
		data[word]++
	}
	ret := make([]string, 0, len(data))
	for k, v := range data {
		if v == 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
*/
