package main

import (
	"fmt"
)

/*
Given a string S and a character C,
return an array of integers representing the shortest distance from the character C in the string.

Example 1:

Input: S = "loveleetcode", C = 'e'
Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]

l ---> e: 3
o --> e: 2
v -> e: 1
e > e: 0
.
.
.
e <---- d -> e: 1 (shortest)
e > e: 0

Note:
S string length is in [1, 10000].
C is a `single` character, and guaranteed to be in string S.
All letters in S and C are lowercase.
*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func shortestToChar(S string, C byte) []int {
	n := len(S)
	pos := -n
	ans := make([]int, len(S))
	for i := 0; i < n; i++ {
		if S[i] == C {
			pos = i
		}
		ans[i] = i - pos
	}
	for i := n - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		ans[i] = min(ans[i], abs(i-pos))
	}
	return ans
}

func main() {
	S := "loveleetcode"
	C := byte('e')
	res := shortestToChar(S, C)
	fmt.Println(res)
}

/*
refer:
-

func shortestToChar(S string, C byte) []int {
	res := []int{}
	for i := range S {
		if S[i] == C {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	}
	count := 9999
	for i, e := range res {
		// 顺着
		if e == 0 {
			count = 0
		} else {
			count++
			res[i] = count
		}
	}
	for i := len(res)-1; i >= 0; i-- {
		// 逆序着
		if res[i] == 0 {
			count = 0
		} else {
			count++
			if res[i] > count {
				res[i] = count
			}
		}
	}
	return res
}
*/
