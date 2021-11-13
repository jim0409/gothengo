package main

/*
Find the minimum length word from a given dictionary words, which has all the letters from the string licensePlate.
Such a word is said to complete the given string `licensePlate`
Here, for letters we ignore case. For example, "P" on the licensePlate still matches "p" on the word.
It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.
The license plate might have the same letter occurring multiple times.
For example, given a licensePlate of "PP", the word "pair" does not complete the licensePlate, but the word "supper" does.

Example 1:
Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
Output: "steps"
Explanation:
The smallest length word that contains the letters "S", "P", "S", and "T".
涵括`licensePlate`的最小字，需要有"S" "P" "S" 以及 "T"
Note that the answer is not "step", because the letter "s" must occur in the word twice.
答案並非`step`是因為，"s"只有出現一次。但實際上要求兩次
Also note that we ignored case for the purposes of comparing whether a letter exists in the word.
(註記:這邊布考慮`licensePlate`不存在於字串中)

Example 2:
Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
Output: "pest"
Explanation:
There are 3 smallest length words that contains the letters "s".
We return the one that occurred first.
有三個符合答案的情境，直接選第一個。

Note:
licensePlate will be a string with length in range [1, 7].
licensePlate will contain digits, spaces, or letters (uppercase or lowercase).
words will have a length in the range [10, 1000].
Every words[i] will consist of lowercase letters, and have length in range [1, 15].
*/

func shortestCompletingWord(licensePlate string, words []string) string {
	var ret string
	src := count(licensePlate) // 回傳一個陣列，該陣列紀錄每一個字元的使用頻率

	// 將單字陣列`words`中的每一單字，逐一取出比較
	for _, word := range words {

		// 非有效字則繼續
		if !isValid(src, word) {
			continue
		}

		// 在單字陣列中，找尋符合條件最小長度單字
		if ret == "" || len(word) < len(ret) {
			ret = word
		}
	}
	return ret
}

// 判斷該單字,`word`是否具有對應字
func isValid(src []int, word string) bool {
	dest := count(word)
	for i := range src {
		if src[i] > dest[i] {
			return false
		}
	}
	return true
}

// 製作一個 count 陣列，轉譯單字長度
func count(word string) []int {
	ret := make([]int, 26)
	for _, c := range word {
		if 'A' <= c && c <= 'Z' {
			ret[c-'A']++
		} else if 'a' <= c && c <= 'z' {
			ret[c-'a']++
		}
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/shortest-completing-word/discuss/370237/Go-golang-0ms-solution

func shortestCompletingWord(licensePlate string, words []string) string {

    s := ""
	minLen := 16
	index := make([]int, 26, 26)
	tmp := strings.ToLower(licensePlate)

	for _, v := range tmp {
		if v >= 'a' && v <= 'z' {
			index[v-'a']++
		}
	}
	for _, v := range words {
		if len(v) >= minLen {
			continue
		}
		if !matches(index, v) {
			continue
		}
		minLen = len(v)
		s = v
	}

	return s

}

func matches(index []int, s string) bool {
	temp := make([]int, 26, 26)
	for _, v := range s {
		temp[v-'a']++
	}
	for i := 0; i < 26; i++ {
		if index[i] > temp[i] {
			return false
		}
	}
	return true
}
*/
