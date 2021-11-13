package main

/*
Given a string which consists of lowercase or uppercase letters,
find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input: "abccccdd"
Output: 7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
*/

func longestPalindrome(s string) int {
	var cnt int
	data := make([]int, 128)
	for _, c := range s {
		data[c]++
		if data[c]%2 == 0 {
			cnt += 2
		}
	}
	if cnt < len(s) {
		cnt++
	}
	return cnt
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/longest-palindrome/discuss/457947/Go-HashSet-solution


func longestPalindrome(s string) int {
	cache := make(map[byte]struct{})
	for i := range s {
		if _, ok := cache[s[i]]; ok {
			delete(cache, s[i])
		} else {
			cache[s[i]] = struct{}{}
		}
	}
	if len(cache) == 0 {
		return len(s)
	}
	return len(s) - len(cache) + 1
}
*/
