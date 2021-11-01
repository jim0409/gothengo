package main

/*
Write a function that reverses a string.
The input string is given as an array of characters char[].

Do not allocate extra space for another array,
you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.


Example 1:
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	// 取總長度的一半
	l := len(s) - 1

	// 前後值互換
	for i := 0; i <= l/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
	return
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/reverse-string/discuss/133322/Go-solution


func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	l := len(s) - 1
	for i := 0; i <= l/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
	return
}
*/
