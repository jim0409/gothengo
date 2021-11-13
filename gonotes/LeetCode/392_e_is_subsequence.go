package main

/*
Given a string s and a string t, check if s is subsequence of t.
You may assume that there is only lower case English letters in both s and t.
t is potentially a very long (length ~= 500,000) string, and s is a short string (<=100).
給定一個字串`s`，以及一個字串t。(假定 s & t 均為小寫英文字母)
且`t`為一個約500,000字的字串；相比之下`s`長度小於100

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters.
(ie, "ace" is a subsequence of "abcde" while "aec" is not).
一個子字串的定義是，如果一個字串是原字串的刪減，但是仍舊保持原始字串的字符(也可以為空字符)
> e.g. ace 為 abcde 的子字串
但 aec 則不是；因為順序錯了

Example 1:
s = "abc", t = "ahbgdc"
Return true.

Example 2:
s = "axc", t = "ahbgdc"
Return false.

Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B,
and you want to check one by one to see if T has its subsequence.
In this scenario, how would you change your code?
*/

// false while execute `s`: "acb"; `t`:"ahbgdc" ... output: true / expect: false
/*
func isSubsequence(s string, t string) bool {
	mT := map[rune]int{}
	for _, j := range t {
		mT[j]++
	}

	for _, j := range s {
		mT[j]--
		if mT[j] < 0 {
			return false
		}
	}
	return true
}
*/

// runtime 24ms; memory:6.8 MB
func isSubsequence(s string, t string) bool {
	// 如果 s 為空字串，直接環傳 true；因為空字串必為任一字串的子集
	if len(s) == 0 {
		return true
	}

	// idx 代表目前 `s`要取的位置， cnt 用來識別最後`len(s)`的長度是否為
	var idx, cnt int
	for i := 0; i < len(t); i++ {
		// i 會一直遞增，去拿取`t`的每一個位置的值，但是`idx`只有當 t[i]==s[idx] 才會做加總
		if t[i] == s[idx] {
			cnt++
			idx++
		}
		// 當 t[i] == s[idx] 時，cnt也會跟著做加總。當`cnt == len(s)` 時，跳出迴圈，返回 true
		if cnt == len(s) {
			return true
		}
	}
	return false
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/is-subsequence/discuss/558140/Go-golang-recursive-8ms-solution


// runtime: 20ms; memory: 6.6MB
func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    if len(t) == 0 {
        return false
    }
    if k := strings.Index(t, string(s[0])); k != -1 {
        return isSubsequence(s[1 :], t[k + 1 :])
    } else {
        return false
    }
}



soution2:
- https://leetcode.com/problems/is-subsequence/discuss/458881/go-clean-code-0ms-beats-100


func isSubsequence(s string, t string) bool {
	var i, j int
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}
*/
