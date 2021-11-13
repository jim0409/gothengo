package main

/*
Given an m * n matrix M initialized with all 0's and several update operations.

Operations are represented by a 2D array,
and each operation is represented by an array with two positive integers a and b,
which means M[i][j] should be added by one for all 0 <= i < a and 0 <= j < b.

You need to count and return the number of maximum integers in the matrix
after performing all the operations.


Example 1:
Input:
m = 3, n = 3
operations = [[2,2],[3,3]]
Output: 4

Explanation:
Initially, M =
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]

After performing [2,2], M =
[[1, 1, 0],
 [1, 1, 0],
 [0, 0, 0]]

After performing [3,3], M =
[[2, 2, 1],
 [2, 2, 1],
 [1, 1, 1]]

So the maximum integer in M is 2, and there are four of it in M. So return 4.

Note:
The range of m and n is [1,40000].
The range of a is [1,m], and the range of b is [1,n].
The range of operations size won't exceed 10,000.
*/

// - https://leetcode.com/problems/range-addition-ii/discuss/382295/Go-golang-2-solutions
func maxCount(m int, n int, ops [][]int) int {
	// 如果 ops 長度為 0 ，拿麼最大值為 0 ，直接回傳 m*n => 0 的總數
	if len(ops) == 0 {
		return m * n
	}
	a := m
	b := n

	for _, v := range ops {
		// 找出`最小值`，`最小值的交集處`即為整個二維方陣的最大值
		if v[0] < a {
			a = v[0]
		}
		if v[1] < b {
			b = v[1]
		}
	}

	return a * b

}

func main() {

}

/*
solution:
- https://leetcode.com/problems/range-addition-ii/discuss/485708/go-clean-code-0ms-beats-100

func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		m, n = int(math.Min(float64(m), float64(op[0]))), int(math.Min(float64(n), float64(op[1])))
	}
	return m * n
}
*/
