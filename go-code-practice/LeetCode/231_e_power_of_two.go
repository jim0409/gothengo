package main


import "fmt"
/*
Given an integer, write a function to determine if it is a power of two.

Example 1:
Input: 1
Output: true 
Explanation: 2^0 = 1

Example 2:
Input: 16
Output: true
Explanation: 2^4 = 16

Example 3:
Input: 218
Output: false
*/


func isPowerOfTwo(n int) bool {
	// 檢查 n 是否為一負數
	if n < 0 {
		return false
	}

	count:=0 // 計算 2 的冪次方個數
	for i := 0; i<32 ;i++{
		if n&1 == 1{
			count++
		}
		if count > 1 {
			return false
		}
		n >>= 1
	}

	return count == 1
}


func main(){
	int1:=6
	fmt.Println(isPowerOfTwo(int1))
}

/*
solution:
- https://leetcode.com/problems/power-of-two/discuss/354239/Go-O(1)-solution


func isPowerOfTwo(n int) bool {
	// Check if n is negative.
	if n < 0 {
		return false
	}

	// If n is the power of two, then it only has one 1 in it's bit representation.
	count := 0
	for i := 0; i < 32; i++ {
		if n&1 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
		n >>= 1
	}
	return count == 1
}
*/