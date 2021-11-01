package main

/*
In MATLAB, there is a very useful function called 'reshape',
which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array,
and two positive integers r and c representing the row number and column number
of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix
in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal,
output the new reshaped matrix; Otherwise, output the original matrix.


Example 1:
Input:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
Output:
[[1,2,3,4]]
Explanation:
The `row-traversing` of nums is [1,2,3,4].
The new reshaped matrix is a 1 * 4 matrix,
fill it row by row by using the previous list.


Example 2:
Input:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
Output:
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix.
So output the original matrix.


Note:
The height and width of the given matrix is in range [1, 100].
The given r and c are all positive.
*/

// https://leetcode.com/problems/reshape-the-matrix/discuss/472194/go-clean-code-4ms-beats-100
func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0]) // 先定義出 row & col 的總長度

	// 檢查總元素個素，倘若不一直接返回
	if r*c != m*n {
		return nums
	}

	// 生成一個新的二維陣列
	ret := make([][]int, r)

	// 生成二維陣列中的子陣列
	for i := range ret {
		ret[i] = make([]int, c)
	}

	// 賦值
	for i := 0; i < r*c; i++ {
		ret[i/c][i%c] = nums[i/n][i%n]
	}

	return ret
}
func main() {

}

/*
solution:
- https://leetcode.com/problems/reshape-the-matrix/discuss/340843/Go-O(n)-with-explanation


func matrixReshape(nums [][]int, nr int, nc int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	// get the row count and column count
	r, c := len(nums), len(nums[0])
	// get the number count
	count := r * c

	// return the origin if the total number is not match
	if count != nr*nc {
		return nums
	}

	// create new matirx
	result := make([][]int, nr)
	for i := 0; i < nr; i++ {
		result[i] = make([]int, nc)
	}

	for i := 0; i < count; i++ {
		oldRow := i / c
		oldCol := i % c
		newRow := i / nc
		newCol := i % nc
		result[newRow][newCol] = nums[oldRow][oldCol]
	}

	return result
}
*/
