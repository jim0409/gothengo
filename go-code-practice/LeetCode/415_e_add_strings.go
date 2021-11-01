package main

/*
Given two non-negative integers num1 and num2 represented as string,
return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

func addStrings(num1 string, num2 string) string {
	ret := make([]byte, len(num1)+1)
	if len(num1) < len(num2) {
		ret = make([]byte, len(num2)+1)
	}
	// carry 表示進位
	// i,j 表示 num1 & num2 的長度
	// k 代表 ret 的長度 (k為一個byte陣列，用來儲存每一個位元的代表數字)
	// 迭代條件為 k > 0 (前面會限定，k為最大的長度+1)
	for carry, i, j, k := byte(0), len(num1)-1, len(num2)-1, len(ret)-1; k >= 0; carry, i, j, k = carry/10, i-1, j-1, k-1 {
		if i >= 0 {
			carry += num1[i] - '0'
		}
		if j >= 0 {
			carry += num2[j] - '0'
		}
		ret[k] = carry%10 + '0'
	}
	if ret[0] == '0' {
		return string(ret[1:])
	}
	return string(ret)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/add-strings/discuss/459554/go-clean-code-0ms-beats-100


func addStrings(num1 string, num2 string) string {
	ret := make([]byte, len(num1)+1)
	if len(num1) < len(num2) {
		ret = make([]byte, len(num2)+1)
	}
	for carry, i, j, k := byte(0), len(num1)-1, len(num2)-1, len(ret)-1; k >= 0; carry, i, j, k = carry/10, i-1, j-1, k-1 {
		if i >= 0 {
			carry += num1[i] - '0'
		}
		if j >= 0 {
			carry += num2[j] - '0'
		}
		ret[k] = carry%10 + '0'
	}
	if ret[0] == '0' {
		return string(ret[1:])
	}
	return string(ret)
}
*/
