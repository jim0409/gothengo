package main

/*
Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.


Example 1:

Input: 6
Output: true
Explanation: 6 = 2 × 3


Example 2:

Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2


Example 3:

Input: 14
Output: false
Explanation: 14 is not ugly since it includes another prime factor 7.

Note:

1 is typically treated as an ugly number.
Input is within the 32-bit signed integer range: [−231,  231 − 1].
*/

// 0ms 2.2MB
func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%5 == 0 {
			num /= 5
		} else if num%3 == 0 {
			num /= 3
		} else {
			return false
		}
	}

	return true
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/ugly-number/discuss/458007/go-8-line-code-0ms-beats-100


// 0ms 2.1MB
func isUgly(num int) bool {
	for i := 2; num > 0 && i < 6; i++ {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1
}

*/
