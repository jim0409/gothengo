package main

/*
Calculate the sum of two integers a and b,
but you are not allowed to use the operator + and -.

Example 1:
Input: a = 1, b = 2
Output: 3

Example 2:
Input: a = -2, b = 3
Output: 1
*/

func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}

	return b
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/sum-of-two-integers/discuss/378190/Go-golang-0ms-solution


func getSum(a int, b int) int {

    for {
        tmp := (a & b) << 1
        a = a ^ b
        b = tmp
        if b == 0 {
            break
        }
    }

    return a

}
*/
