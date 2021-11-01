package main

import (
	"fmt"
	"strings"
)

/*
Given two strings A and B,
find the minimum number of times A has to be repeated such that B is a substring of it.
If no such solution, return -1.

Example1:
A = "abcd"
B = "cdabcdab"
Return 3

Explanation:
because by repeating A three times (“abcdabcdabcd”), B is a substring of it;
and B is not a substring of A repeated two times ("abcdabcd").

Note:
The length of A and B will be between 1 and 10000.
*/

func repeatedStringMatch(A string, B string) int {
	res := 1
	for res <= len(B)/len(A)+2 {
		if strings.Contains(strings.Repeat(A, res), B) {
			return res
		}
		res++
	}
	return -1
}

// func repeatedStringMatch(A string, B string) int {
// 	res := 1
// 	for res <= len(B)/len(A)+2 {
// 		if containStr(repStr(A, res), B) {
// 			return res
// 		}
// 		res++
// 	}
// 	return -1
// }

// func containStr(s string, sub string) bool {
// 	if s == sub {
// 		return true
// 	}

// 	subl := len(sub)
// 	for i := 0; i < len(s)-subl+1; i++ {
// 		if s[i:i+subl] == sub {
// 			return true
// 		}
// 	}

// 	return false
// }

// func repStr(s string, t int) string {

// 	var res string
// 	for i := 0; i < t; i++ {
// 		res += s
// 	}
// 	return res
// }

func main() {

	// str1 := "abcdabcdabcd"
	str1 := "abcd"
	str2 := "cdabcdab"

	fmt.Println(containStr(str1, str2))
	fmt.Println(containStr(repStr(str1, 3), str2))

	str3 := "a"
	str4 := "aa"
	fmt.Println(containStr(repStr(str3, 2), str4))

	str5 := "aaac"
	str6 := "aac"
	fmt.Println(containStr(repStr(str5, 1), str6))
}

/*
solution:
- https://leetcode.com/problems/repeated-string-match/discuss/493380/go-clean-code-0ms-beats-100

func repeatedStringMatch(A string, B string) int {
	cnt := int(math.Ceil(float64(len(B)) / float64(len(A))))
	rpt := strings.Repeat(A, cnt)
	if strings.Contains(rpt, B) {
		return cnt
	}
	if strings.Contains(rpt+A, B) {
		return cnt + 1
	}
	return -1
}
*/

/*
solution:
- https://leetcode.com/problems/repeated-string-match/discuss/421935/Go-golang-solutions

func repeatedStringMatch(A string, B string) int {
    res := 1
    for res <= len(B) / len(A) + 2 {
        if strings.Contains(strings.Repeat(A, res), B) { return res }
        res++
    }
    return -1
}
*/
