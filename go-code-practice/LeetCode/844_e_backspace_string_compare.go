package main

import "fmt"

/*
Given two strings S and T, return if they are equal when both are typed into empty text editors.
`#`` means a backspace character.
Note that after backspacing an empty text, the text will continue empty.

Example 1:
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".

Example 2:
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".

Example 3:
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".

Example 4:
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".

Note:
1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:
Can you solve it in O(N) time and O(1) space?
*/
func convertWords(s string) string {
	bs := byte('#')
	var words []byte
	for _, j := range []byte(s) {
		if j == bs && len(words) >= 1 {
			words = words[:len(words)-1]
			continue
		}
		if j != bs {
			words = append(words, j)
		}
	}
	return string(words)
}

func backspaceCompare(S string, T string) bool {
	return convertWords(S) == convertWords(T)
}

func main() {
	S1 := "ab##"
	T1 := "c#d#"
	fmt.Println(backspaceCompare(S1, T1))

	S2 := "a##c"
	T2 := "#a#c"
	fmt.Println(backspaceCompare(S2, T2))

	S3 := "y#fo##f"
	T3 := "y#f#o##f"
	fmt.Println(backspaceCompare(S3, T3))
}

/*
refer:
- https://leetcode.com/problems/backspace-string-compare/discuss/307072/go

func backspaceCompare(S string, T string) bool {
	return dec(S) == dec(T)
}

func dec(T string) string {
	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			if i > 0 {
				T = T[:i-1] + T[i+1:]
				i -= 2
			} else {
				T = T[:i] + T[i+1:]
				i--
			}
		}
	}
	return T
}
*/
