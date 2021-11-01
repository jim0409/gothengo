package main

/*
Given a paragraph and a list of banned words,
return the most frequent word that is not in the list of banned words.

It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.

Words in the paragraph are not case sensitive.

The answer is in lowercase.


Example:
Input:
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.


Note:
1 <= paragraph.length <= 1000.
0 <= banned.length <= 100.
1 <= banned[i].length <= 10.
The answer is unique, and written in lowercase
(even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
There are no hyphens or hyphenated words.
Words only consist of letters, never apostrophes or other punctuation symbols.
*/

func mostCommonWord(paragraph string, banned []string) string {
	// record banned words that is not more than 1000
	bannedCache := make(map[string]struct{})
	for _, b := range banned {
		bannedCache[b] = struct{}{}
	}

	// declare words for lowercase transfer
	// also construct sentence to string array
	words := make([]string, 0)
	var tmp string
	for _, c := range paragraph {
		if c >= 'A' && c <= 'Z' {
			tmp += string(c + 32)
		} else if c >= 'a' && c <= 'z' {
			tmp += string(c)
		} else {
			if tmp != "" {
				words = append(words, tmp)
			}
			tmp = ""
		}
	}
	// append the last word
	if tmp != "" {
		words = append(words, tmp)
	}

	// count words frequency
	cache := make(map[string]int)
	for _, w := range words {
		cache[w]++
	}

	max, maxWord := 0, ""
	for word, freq := range cache {
		if _, ok := bannedCache[word]; ok || (max > 0 && freq < max) {
			continue
		}
		max = freq
		maxWord = word
	}
	return maxWord
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/most-common-word/discuss/580114/go-clean-code-0ms-beats-100


func mostCommonWord(paragraph string, banned []string) string {
	data := make(map[string]int)
	for _, word := range banned {
		data[word] = math.MinInt32
	}
	var ret string
	for max, i, j := 0, 0, 0; i < len(paragraph) && j < len(paragraph); {
		for i = j; i < len(paragraph) && !isChar(paragraph[i]); i++ {
		}
		for j = i; j < len(paragraph) && isChar(paragraph[j]); j++ {
		}
		if i >= len(paragraph) {
			break
		}
		word := strings.ToLower(paragraph[i:j])
		data[word]++
		if data[word] > max {
			ret = word
			max = data[word]
		}
	}
	return ret
}

func isChar(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}
*/
