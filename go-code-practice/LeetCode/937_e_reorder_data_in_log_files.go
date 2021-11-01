package main

import (
	"sort"
	"strings"
)

/*
You are given an array of logs.
Each log is a space-delimited string of words, where the first word is the identifier.

There are two types of logs:

Letter-logs: All words (except the identifier) consist of lowercase English letters.
Digit-logs: All words (except the identifier) consist of digits.

Reorder these logs so that:
1. The letter-logs come before all digit-logs.
2. The letter-logs are sorted lexicographically by their contents. If their contents are the same, then sort them lexicographically by their identifiers.
3. The digit-logs maintain their relative ordering.
Return the final order of the logs.


Example 1:
Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
Explanation:
The letter-log contents are all different, so their ordering is "art can", "art zero", "own kit dig".
The digit-logs have a relative order of "dig1 8 1 5 1", "dig2 3 6".

Example 2:
Input: logs = ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]


Constraints:
1 <= logs.length <= 100
3 <= logs[i].length <= 100
All the tokens of logs[i] are separated by a single space.
logs[i] is guaranteed to have an identifier and at least one word after the identifier.
*/

// https://leetcode.com/problems/reorder-data-in-log-files/discuss/506848/Go-Stable-sort-solution
// https://blog.csdn.net/fuxuemingzhu/article/details/83961188
/*
題意:
每條日誌是由空格隔開的一些字符串，第一個字符串是標誌符。

日誌有兩種格式(標誌符不在以下討論範圍中)
1. 文字型日誌，皆由英文小寫組成
2. 數字型日誌，皆由數字組成

需要重新排序，排序規則如下
1. 文字型日誌皆排列在數字型日誌前
2. 文字型日誌如果重複則需要排序，文字型日誌排序順序為根據其 第一字(標誌符) 為準
3. 數字型日誌在排序後只要維持相對位置即可

e.g.
logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
dig1, dig2, let1, let2, let3 皆為標誌符
其中 "let1 arc can", "let2 own kit dig" 及 "let3 ar zero" 為文字型日誌
排序結果為 "art can", "art zero" 及 "own kit dig"，再將數字型日誌放置在文字型日誌後做相對位置排序

Output: ["let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"]

... 題意很難理解的一題...

1.將文字型日誌與數字型日誌分開儲存
2.對文字型的內容做排序
3.排序過後再將數字型日誌接回
*/
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		firstSpace, secondSpace := strings.Index(logs[i], " "), strings.Index(logs[j], " ")
		firstChar, secondChar := logs[i][firstSpace+1], logs[j][secondSpace+1]
		if firstChar <= '9' {
			if secondChar <= '9' {
				return i < j
			}
			return false
		}
		if secondChar <= '9' {
			return true
		}
		comparison := strings.Compare(logs[i][firstSpace+1:], logs[j][secondSpace+1:])
		if comparison == 0 {
			comparison = strings.Compare(logs[i][:firstSpace], logs[j][:secondSpace])
		}
		return comparison < 0
	})
	return logs
}

func main() {}

/*
refer:
- https://leetcode.com/problems/reorder-data-in-log-files/discuss/581910/go-clean-code-4ms-beats-100

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		k, l := strings.Index(logs[i], " "), strings.Index(logs[j], " ")
		switch {
		case logs[i][k+1] <= '9':
			return false
		case logs[j][l+1] <= '9':
			return true
		default:
			if logs[i][k+1:] == logs[j][l+1:] {
				return logs[i][:k] < logs[j][:l]
			}
			return logs[i][k+1:] < logs[j][l+1:]
		}
	})
	return logs
}
*/
