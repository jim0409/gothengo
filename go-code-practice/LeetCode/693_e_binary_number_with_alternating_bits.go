package main

/*
Given a positive integer, check whether it has alternating bits:
namely, if two adjacent bits will always have different values.

Example 1:
Input: 5
Output: True
Explanation:
The binary representation of 5 is: 101

Example 2:
Input: 7
Output: False
Explanation:
The binary representation of 7 is: 111.

Example 3:
Input: 11
Output: False
Explanation:
The binary representation of 11 is: 1011.

Example 4:
Input: 10
Output: True
Explanation:
The binary representation of 10 is: 1010.
*/

func hasAlternatingBits(n int) bool {
	n ^= n >> 1
	return n&(n+1) == 0
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/binary-number-with-alternating-bits/discuss/315061/Go-100-O(n)-simple-and-concise

func hasAlternatingBits(n int) bool {
	s := 1
	for s <= n {
		s <<= 1
	}
	s -= 1

	return (n >> 1) ^ s == n
}
*/
