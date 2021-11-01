package main

import (
	"bytes"
	"fmt"
)

/*
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match,
such that there is a bijection between a letter in pattern and
a non-empty word in str.

Example 1:
Input: pattern = "abba", str = "dog cat cat dog"
Output: true

Example 2:
Input:pattern = "abba", str = "dog cat cat fish"
Output: false

Example 3:
Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false

Example 4:
Input: pattern = "abba", str = "dog dog dog dog"
Output: false

Notes:
You may assume `pattern` contains only lowercase letters,
and `str` contains lowercase letters that may be separated by a single space.
*/

// 實作 strings.Split ...
func splitSpace(str string) []string {
	strArr := []string{}
	var tmpBuff bytes.Buffer
	for _, j := range str {
		if j == ' ' {
			strArr = append(strArr, tmpBuff.String())
			tmpBuff = bytes.Buffer{}
			continue
		}
		tmpBuff.WriteRune(j)
	}
	strArr = append(strArr, tmpBuff.String())

	return strArr
}

// pattern => ['r','u','n','e']
// str => ['dog','dog','dog','dog']
func wordPattern(pattern string, str string) bool {
	list := splitSpace(str)
	dict1 := make(map[byte]string)
	dict2 := make(map[string]byte)
	if len(pattern) != len(list) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if v, ok := dict1[pattern[i]]; !ok {
			dict1[pattern[i]] = list[i]
		} else {
			if v != list[i] {
				return false
			}
		}

		if w, ok := dict2[list[i]]; !ok {
			dict2[list[i]] = pattern[i]
		} else {
			if w != pattern[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))
}

/*
solution:
- https://leetcode.com/problems/word-pattern/discuss/458086/go-19-line-code-0ms-beats-100


func wordPattern(pattern string, str string) bool {
	slice := strings.Fields(str)
	if len(pattern) != len(slice) {
		return false
	}
	data1, data2 := make(map[byte]int, len(pattern)), make(map[string]int, len(slice))
	for i := 0; i < len(pattern); i++ {
		if _, ok := data1[pattern[i]]; !ok {
			data1[pattern[i]] = i
		}
		if _, ok := data2[slice[i]]; !ok {
			data2[slice[i]] = i
		}
		if data1[pattern[i]] != data2[slice[i]] {
			return false
		}
	}
	return true
}


solution2:
- https://leetcode.com/problems/word-pattern/discuss/415053/Go-golang-0ms-solutions

func wordPattern(pattern string, str string) bool {
    s := strings.Fields(str)
    if len(pattern) != len(s) { return false }
    for i := 0; i < len(s); i++ {
        for j := i+1; j < len(s); j++ {
            if pattern[i] == pattern[j] && s[i] != s[j] || pattern[i] != pattern[j] && s[i] == s[j] {
                return false
            }
        }
    }
    return true
}
*/
