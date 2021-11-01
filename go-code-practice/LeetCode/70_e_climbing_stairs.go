package main

import "fmt"

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

*/

// solution: https://leetcode.com/problems/climbing-stairs/discuss/351095/Go-DP-solution
func climbStairs(n int) int {
	// 當台階小於三階的時候，有的排列組合剛好為n階種
	if n < 3 {
		return n
	}

	one, two := 1, 1
	// if n=4;
	// [i=2]next=2, one=1, two=2 --> [i=3]next=3, one=2, two=3
	for i := 2; i < n; i++ {
		next := one + two
		one = two
		two = next
	}
	return one + two
}

func main() {
	input1 := 4
	fmt.Println(climbStairs(input1))
}

/*
Fibonacci Go Solution:
- https://leetcode.com/problems/climbing-stairs/discuss/402652/Fibonacci-GO-Solution

func climbStairs(n int) int {
    if n <= 1 {
        return 1
    }

    oneStep := 1
    twoStep := 1

    for i:= 2; i<n; i++ {
        oneStep, twoStep = oneStep + twoStep, oneStep
    }

    return oneStep + twoStep
}

*/

/*
Hint:
To reach nth step, what could have been your previous steps? (Think about the step sizes)

費波曼數列
2,3,5,8,13,...

n=3

111
12
21

n=4

111 1
12  1
21  1
--
11 2
22 (整除會多這項)

n=5

1111 1
121 1
211 1
112 1
22 1
--
12 2
21 2
111 2
*/
