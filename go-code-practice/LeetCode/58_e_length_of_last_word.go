package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Given a string s consists of upper/lower-case alphabets

and empty space characters ' ',

return the length of last word

(last word means the last appearing word if we loop from left to right)

in the string.

If the last word does not exist, return 0.

Note: A word is defined as a maximal substring consisting

of non-space characters only.
*/

// any chance to make a O(n) instead of O(2n) ??
func lengthOfLastWord(s string) int {
	// 先把最後一位數抓出來
	pos := len(s) - 1
	// edge case 防止， "a  "的狀況。所以要先把a以後的空白都移除
	for ; pos >= 0; pos-- {
		if !unicode.IsSpace(rune(s[pos])) {
			break
		}
	}

	// 判斷字串的長度
	s = s[:pos+1]
	if len(s) == 0 {
		return 0
	}

	// 從後面數過來，當遇到空白時。可以斷定最後一位數的pos是從哪邊開始
	// "  a" 可以得到一個字串長度為三的數字。但是a前面的空白所在的位數為1，所以a就是s[1:2]
	for ; pos >= 0; pos-- {
		if unicode.IsSpace(rune(s[pos])) {
			break
		}
	}

	return len(s[pos+1:])
}

func main() {
	str1 := "Hello World"
	fmt.Println(lengthOfLastWord(str1))

	str2 := "H"
	fmt.Println(lengthOfLastWord(str2))

	str3 := " "
	fmt.Println(lengthOfLastWord(str3))

	str4 := "Today is a nice day"
	fmt.Println(lengthOfLastWord(str4))

	str5 := " a"
	fmt.Println(lengthOfLastWord(str5))

	strings.Split("123", "123")

}

/*
soulution:
1.You should remove all space char of tail of the string.
2.Well, you count the last length of the word.It's better to count the position of the word from the tail of the string.
3.That's all.


- https://leetcode.com/problems/length-of-last-word/discuss/21924/Fast-Code-Of-Go-Go-Go

func lengthOfLastWord(s string) int {
	pos := len(s) - 1
	for ; pos >= 0; pos-- {
		if !unicode.IsSpace(rune(s[pos])) {
			break
		}
	}
	s = s[:pos + 1]
	if len(s) == 0 {
		return 0
	}

	for ; pos >= 0; pos-- {
		if unicode.IsSpace(rune(s[pos])) {
			break
		}
	}
	return len(s[pos+1:])
}
*/
