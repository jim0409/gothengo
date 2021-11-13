package main

/*
Given a non-empty string s, you may delete at most one character.
Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True

Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.

Note:
The string will only contain lowercase characters a-z.
The maximum length of the string is 50000.
*/

// https://leetcode.com/problems/valid-palindrome-ii/discuss/493271/go-clean-code-8ms-beats-100
func validPalindrome(s string) bool {
	// 左右同時逼近 (i) -->  <---(j)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
		}
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func main() {

}

/*
solution:
- https://leetcode.com/problems/valid-palindrome-ii/discuss/507206/Go-12ms-100-clear-solution


// 將三種可能會為回文的情境條列出來~
func validPalindrome(s string) bool {
    valid, l, r := isPalindrome(s, 0, len(s)-1)
    if valid {
        return true
    }
    if valid, _, _ := isPalindrome(s, l+1, r); valid {
        return true
    }
    if valid, _, _ := isPalindrome(s, l, r-1); valid {
        return true
    }
    return false
}

func isPalindrome(s string, l, r int) (bool, int, int) {
    for l < r {
        if s[l] != s[r] {
            return false, l, r
        }
        l++
        r--
    }
    return true, 0, 0
}
*/
