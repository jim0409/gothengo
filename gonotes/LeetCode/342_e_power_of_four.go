package main

/*
Given an integer (signed 32 bits),
write a function to check whether it is a power of 4.

Example 1:
Input: 16
Output: true

Example 2:
Input: 5
Output: false

Follow up: Could you solve it without loops/recursion?
*/

func isPowerOfFour(num int) bool {
	return num != 0 && num&(num-1) == 0 && num&1431655765 == num
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/power-of-four/discuss/236782/go-wtih-explain


二進制中
1 -> 1
4 -> 100
16 -> 10000
64 -> 100000
...
...
1 都在 奇數位，且均為1

// refer:
// - https://stackoverflow.com/a/6260496
// 2^32 - 1
num&1431655765==num 保證 1 都在奇數位
num&(num-1)==0 保證只有一個 1
綜上，就只有一個奇數位的 1 如 100 是 4
*/

/*
solution:
- https://leetcode.com/problems/power-of-four/discuss/406021/Go-golang-0ms-solution

func isPowerOfFour(num int) bool {
    i := 1
    for i < num { i *= 4 }
    return i == num
}
*/
