package main

/*
Given a 2D integer array matrix, return the transpose of matrix.

The transpose of a matrix is the matrix flipped over its main diagonal,
switching the matrix's row and column indices.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]

Example 2:
Input: matrix = [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]


Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 1000
1 <= m * n <= 105
-109 <= matrix[i][j] <= 109
*/

func transpose(A [][]int) [][]int {
	ret := make([][]int, len(A[0]))
	for i := range ret {
		ret[i] = make([]int, len(A))
		for j := range ret[i] {
			ret[i][j] = A[j][i]
		}
	}
	return ret
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/transpose-matrix/discuss/149428/32-ms-Go-solution

func transpose(A [][]int) [][]int {
    n := len(A)
    m := len(A[0])
    T := make([][]int, m)

    for i := 0; i < m; i++ {
        T[i] = make([]int, n)
    }

    for i, Ai := range A {
        for j, v := range Ai {
            T[j][i] = v
        }
    }

    return T
}
*/
