package main

import (
	"fmt"
	"strings"
)

/*
Given a non-empty string check if it can be constructed by taking a substring of it
and appending multiple copies of the substring together.
考慮給定的字串中，是否含有子字串

You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.


Example 1:
Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.

Example 2:
Input: "aba"
Output: False

Example 3:
Input: "abcabcabcabc"
Output: True
Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
*/

func repeatedSubstringPattern(s string) bool {
	// 如果 len(s) 長度為 0 代表不存在 substring
	if len(s) == 0 {
		return false
	}

	size := len(s)
	// 1:2n-1 表示 取 `兩倍字串`的 總長度`的 去頭去尾
	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}

func main() {
	input1 := "abab"
	fmt.Println(repeatedSubstringPattern(input1))

	input2 := "abcabcabcabc"
	fmt.Println(repeatedSubstringPattern(input2))

	input3 := "aba"
	fmt.Println(repeatedSubstringPattern(input3))
}

/*
solution:
- https://leetcode.com/problems/repeated-substring-pattern/discuss/555811/GO-Substring-solution


func repeatedSubstringPattern(s string) bool {
    if len(s) == 0 {
        return false
    }

    size := len(s)
    ss := (s+s)[1:size*2-1]

    return strings.Contains(ss,s)
}
*/

/*
// 最快的解答
solution2:
- https://leetcode.com/problems/repeated-substring-pattern/discuss/94354/Fastest-go-solution-so-far-(12ms)


// 1. 先找出該字串的最後一個字元，其中該字元與第一個字元相同
// 2. 從起始的群集開始到最後一個群集，檢查所有的符號。看兩個群集是否相同
//		=> 從頭開始，檢查到最後一個群集。只要失敗的話，就回到第一步重新劃分群集

func repeatedSubstringPattern(s string) bool {
	// 如果字串長度小於 2 直接回傳錯誤
    if len(s) < 2 {
        return false
    }

	// 讓 first 為第一個字元，max為該字串長度-1
    first := s[0]
    max := len(s) - 1

	// 找出與第一個 字元 相同 的 字元的位置
	dupIndex := findPrevIndex(first, s, max)

	// 如果 dupIndex 不為 0
    for dupIndex > 0 {

		// step 表示要查找的次數
        step := max - dupIndex + 1
        if step > 0 && len(s) % step == 0 {
			// ok: 存在, from: 出現相同字元的位置, to:該字串長度大小-1
            ok := true
            from, to := dupIndex, max
            for ok && from > 0 {
				//
                ok = checkFromTo(s, from, to)
                from -= step
                to -= step
            }
            if ok {
                return true
            }
        }

        dupIndex = findPrevIndex(first, s, max - step)
    }

    return false
}


// 檢查 兩群字元組 是否均為一致  `s[0:to-from]  vs  s[from:to]`
func checkFromTo(s string, from, to int) bool {
    var i int
    for i = 0; i <= to - from; i++ {
        if s[i] != s[from + i] {
            return false
        }
    }
    return true
}


// 檢查 字串 中，是否存在與該 字元 相同的值，並且回傳 位址
func findPrevIndex(ch byte, s string, max int) int {
    for i := max; i > 0; i-- {
        if s[i] == ch {
            return i
        }
    }
    return 0
}
*/
