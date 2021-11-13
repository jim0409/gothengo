package main

import "fmt"

/*
Given an integer, write an algorithm to convert it to hexadecimal.
For negative integer, two’s complement method is used.
// https://en.wikipedia.org/wiki/Two%27s_complement

Note:

All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s.
If the number is zero, it is represented by a single zero character '0';
otherwise, the first character in the hexadecimal string will not be the zero character.

The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.

Example 1:

Input: 26
Output: "1a"

Example 2:

Input: -1
Output: "ffffffff"
*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var n uint64
	if num < 0 {
		n = 1<<32 + uint64(num)
	} else {
		n = uint64(num)
	}
	hex := "0123456789abcdef"
	var s string
	for n > 0 {
		a := n % 16
		fmt.Println(string(hex[a]))
		s = string(hex[a]) + s
		n /= 16
	}
	return s
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/convert-a-number-to-hexadecimal/discuss/327397/Go-easy-to-understand


func toHex(num int) string {

	if num == 0 {
		return "0"
	}

	hash := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	result := ""
	for num != 0 {
		result = string(hash[num&15]) + result
		num = int(uint32(num) >> 4)   //Right shift of unsigned bit

	}
	return result

}
*/
