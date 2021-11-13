package main

import "fmt"

// Valid Parentheses
/*
Given a string containing just the characters

'(', ')', '{', '}', '[' and ']',

determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
*/

func isValid(s string) bool {
	// 宣告一個 stack 來儲存 rune 的陣列，做前後比對
	stack := []rune{}

	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		// 因為字串是先從 '(' . '{' 及 '[' 開始，所以前面一定都對不到
		// 前面主要在做儲存stack的動作，後面再來比較
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		// 當有比較到的時候，stack最後一位pop out; v=')', 則 m[v]='('
		if m[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}

func main() {
	for i, j := range m {
		fmt.Println(i, j)
	}

}
