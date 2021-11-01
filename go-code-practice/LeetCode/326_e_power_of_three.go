package main

/*
Given an integer, write a function to determine if it is a power of three.

Example 1:
Input: 27
Output: true

Example 2:
Input: 0
Output: false

Example 3:
Input: 9
Output: true

Example 4:
Input: 45
Output: false

Follow up:
Could you do it without using any loop / recursion?
*/

func isPowerOfThree(n int) bool {
	if (n != 1) && (n == 0 || n%3 != 0) {
		return false
	}

	// 一直除，除到 n == 0
	for n != 0 {
		n = n / 3
		// 如果 n/3 的餘數不為 0 且 n 本身不為 1 return false
		if n%3 != 0 && n != 1 {
			return false
		}
	}

	return true
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/power-of-three/discuss/460392/go-solution-without-loop


func isPowerOfThree(n int) bool {
    if n <= 0 {
        return false
	}
	// meta numbers
	// - https://metanumbers.com/4052555153018976267
	return 4052555153018976267 % n == 0
}


solution2:
- https://leetcode.com/problems/power-of-three/discuss/402177/Go-golang-clean-solution


func isPowerOfThree(n int) bool {
    a := 1
    for a < n {
        a *= 3
    }
    return a == n
}
*/
