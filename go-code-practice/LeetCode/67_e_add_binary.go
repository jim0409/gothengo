package main

import (
	"fmt"
	"strconv"
)

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"

Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

func addBinary(a string, b string) string {
	// 用來儲存最後的數字
	var result string
	var carry bool
	// 因為a, b 都是字串，且題目沒有明示哪一個字串長度比較長，所以要比較道兩邊字串都結束才能停止
	// 定義 i 及 j 來依序拿取字串的每個元素
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		sum := 0
		// 當 i 及 j 仍舊大於等於 0，表示還沒取到第一位需要繼續往下取
		if i >= 0 && j >= 0 {
			// 如果 a 當下的位數為 '1' 則對 sum++
			if a[i] == '1' {
				sum++
			}
			if b[j] == '1' {
				sum++
			}
			i--
			j--

			// 要分開比，因為 a & b 的長度不一定相同
		} else if i >= 0 {
			if a[i] == '1' {
				sum++
			}
			i--
		} else {
			if b[j] == '1' {
				sum++
			}
			j--
		}
		if carry {
			sum++
		}
		result = string(strconv.Itoa(sum%2)) + result
		carry = sum > 1
	}
	if carry {
		result = "1" + result
	}
	return result
}

func main() {
	a1 := "11"
	b1 := "1"
	fmt.Println(addBinary(a1, b1))

}

/*
solution:
- https://leetcode.com/problems/add-binary/discuss/344762/Go-Iterative-solution


func addBinary(a string, b string) string {
	var result string
	var carry bool
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		sum := 0
		if i >= 0 && j >= 0 {
			if a[i] == '1' {
				sum++
			}
			if b[j] == '1' {
				sum++
			}
			i--
			j--
		} else if i >= 0 {
			if a[i] == '1' {
				sum++
			}
			i--
		} else {
			if b[j] == '1' {
				sum++
			}
			j--
		}
		if carry {
			sum++
		}
		result = string(strconv.Itoa(sum%2)) + result
		carry = sum > 1
	}
	if carry {
		result = "1" + result
	}
	return result
}
*/
