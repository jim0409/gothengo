package main

import "fmt"

func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		max = maxPalindrome(s, i, i, max)
		max = maxPalindrome(s, i, i+1, max)
	}
	return max
}

func maxPalindrome(s string, i, j int, max string) string {
	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(max) < len(sub) {
		return sub
	}
	return max
}

func main() {
	// s1 := "babad" // expect output: bab or aba
	// fmt.Println(longestPalindrome(s1))

	s2 := "aa"
	fmt.Println(longestPalindrome(s2))

}

// before 2020/2/16
/*
func longestPalindrome(s string) string {
	var mapS = make(map[string]int)
	// var chooseMap  []string{}
	var maxCount = 0
	for _, j := range s {
		if string(j) != "" {
			mapS[string(j)]++
		}
	}

	// calculate max count
	for _, j := range mapS {
		if maxCount < j {
			maxCount = j
		}
	}

	for i, j := range mapS {
		if j == maxCount {
			return i
		}
	}
	return "conflict"
}
*/

/*
// 參考解
func longestPalindrome(s string) string {
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}

	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1

	// 在 for 循环中
	// b 代表回文的**首**字符索引号，
	// e 代表回文的**尾**字符索引号，
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让b，e分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让b，e分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < len(s); { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		if len(s)-i <= maxLen/2 {
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了。
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// 循环结束后，s[b:e+1]是一串相同的字符串
		}

		// 下一个回文的`正中间段`的首字符只会是s[e+1]
		// 为下一次循环做准备
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// 循环结束后，s[b:e+1]是这次能找到的最长回文。
		}

		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}

*/
