package main

/*
Given a List of words, return the words that can be typed using letters of alphabet on only one row's
of American keyboard like the image below.
給定一個字串陣列，將陣列中的字串。可以在鍵盤上的同一列回傳出來的字串回傳出來

Example:

Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]
*/

func findAlphabet(s string) bool {

	mp := map[rune]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}

	row := 0
	for k, v := range s {

		// 轉換大小寫
		if v >= 65 && v <= 90 {
			v = v + 32
		}

		// 初始值的row定義出mp[v]，之後的mp[v]如果撈不到值。表示不是同一列的
		if k == 0 {
			row = mp[v]
		}

		// 逐字元篩選是否為同一列的
		if v, ok := mp[v]; ok && row != v {
			return false
		}
	}
	return true
}
func findWords(words []string) []string {

	var ans []string
	for _, v := range words {
		if findAlphabet(v) {
			ans = append(ans, v)
		}
	}
	return ans
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/keyboard-row/discuss/468003/go-clean-code-0ms-beats-100


func findWords(words []string) []string {
	ret := make([]string, 0, len(words))
	for _, word := range words {
		if isValid(word) {
			ret = append(ret, word)
		}
	}
	return ret
}

func isValid(word string) bool {
	var data = []byte{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}
	for i := 1; i < len(word); i++ {
		if data[toLower(word[i])-'a'] != data[toLower(word[0])-'a'] {
			return false
		}
	}
	return true
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}
*/
