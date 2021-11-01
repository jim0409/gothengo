package main

import "strings"

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.


Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true


Example 2:

Input: "race a car"
Output: false
*/

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(a byte) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}
func main() {

}

/*
solution with regexp :
https://leetcode.com/problems/valid-palindrome/discuss/235636/Go-16ms

func isPalindrome(s string) bool {

    reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
    processedString := reg.ReplaceAllString(s,"")
    processedString = strings.ToLower(processedString)
    for i:=0; i<len(processedString)/2; i++ {
        if processedString[i] != processedString[len(processedString)-1-i]{
            return false
        }
    }
    return true
}
*/

/*
solution without regexp
https://leetcode.com/problems/valid-palindrome/discuss/262475/go-solution

func isPalindrome(s string) bool {
	if len(s) <= 0 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1

	for i <= j {
		if isValid(s[i]) && isValid(s[j]) && (s[i] == s[j]) {
			i++
			j--
            continue
		} else if isValid(s[i]) && isValid(s[j]) && (s[i] != s[j]) {
			return false
		}
		if !isValid(s[i]) {
			i++
		}
		if !isValid(s[j]) {
			j--
		}
	}
	return true
}

func isValid(b byte) bool {
	if ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}
*/
