package main

import (
	"fmt"
)

/*
Given two strings s and t ,
write a function to determine if t is an anagram of s.
< anagram: 字謎 > : 易位構詞~
https://en.wikipedia.org/wiki/Anagram

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters?
How would you adapt your solution to such case?
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}
	for _, j := range s {
		m[j]++
	}

	for _, j := range t {
		m[j]--
	}

	for i, _ := range m {
		if m[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	input1 := "anagram"
	input2 := "nagaram"
	fmt.Println(isAnagram(input1, input2))

	input3 := "a"
	input4 := "b"
	fmt.Println(isAnagram(input3, input4))

	input5 := "aacc"
	input6 := "ccac"
	fmt.Println(isAnagram(input5, input6))
}

/*
solution:
- https://leetcode.com/problems/valid-anagram/discuss/260542/go-solution


func isAnagram(s string, t string) bool {
    if len(s) == 0 && len(t) == 0 {
        return true
    }
    if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
        return false
    }

    buf := make([]int, 26)
    for _, v := range s {
        buf[v-'a']++
    }
    for _, v := range t {
        buf[v-'a']--
    }

    for _, v := range buf {
        if v != 0 {
            return false
        }
    }
    return true
}
*/
