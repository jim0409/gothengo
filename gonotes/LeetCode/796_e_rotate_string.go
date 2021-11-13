package main

/*
We are given two strings, A and B.
A shift on A consists of taking string A and moving the leftmost character to the rightmost position.
給定兩組字串(A、B)，判斷字串A做了一定位移後是否恰好可以為B字串

For example, if A = 'abcde', then it will be 'bcdea' after one shift on A.
Return True if and only if A can become B after some number of shifts on A.


Example 1:
Input: A = 'abcde', B = 'cdeab'
Output: true


Example 2:
Input: A = 'abcde', B = 'abced'
Output: false


Note:
A and B will have length at most 100.
*/

func rotateString(A string, B string) bool {

	if len(A) == 0 && len(B) == 0 {
		return true
	}

	newString := A + A

	for i := 0; i < len(newString)-len(A); i++ {
		if newString[i:len(A)+i] == B {
			return true
		}
	}
	return false
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/rotate-string/discuss/579350/go-clean-code-0ms-beats-100

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
*/
