package main

/*
You are given an n x n grid where we place some 1 x 1 x 1 cubes that are axis-aligned with the x, y, and z axes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of the cell (i, j).

We view the projection of these cubes onto the xy, yz, and zx planes.

A projection is like a shadow, that maps our 3-dimensional figure to a 2-dimensional plane.

We are viewing the "shadow" when looking at the cubes from the top, the front, and the side.

Return the total area of all three projections.

Example 1:
Input: grid = [[1,2],[3,4]]
Output: 17
Explanation: Here are the three projections ("shadows") of the shape made with each axis-aligned plane.
// x投影 => 10(3*4 - 1*2)
// y投影 => 3
// z投影 => 4

Example 2:
Input: grid = [[2]]
Output: 5
// x投影 => 2
// y投影 => 2
// z投影 => 1

Example 3:
Input: grid = [[1,0],[0,2]]
Output: 8

Example 4:
Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
Output: 14

Example 5:
Input: grid = [[2,2,2],[2,1,2],[2,2,2]]
Output: 21


Constraints:
n == grid.length
n == grid[i].length
1 <= n <= 50
0 <= grid[i][j] <= 50
*/

// https://leetcode.com/problems/projection-area-of-3d-shapes/discuss/347540/Go-golang-4ms-clean-solution
// 目的: 找到投影面積總和
func projectionArea(grid [][]int) int {

	count := 0

	for i, v := range grid {

		col := 0
		row := 0

		for I, V := range v {
			if V > 0 {
				count++
			}
			if V > col {
				col = V
			}

			if grid[I][i] > row {
				row = grid[I][i]
			}
		}
		count += row
		count += col
	}
	return count
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/projection-area-of-3d-shapes/discuss/581628/go-clean-code-4ms-beats-100

func projectionArea(grid [][]int) int {
	var ret int
	for i := range grid {
		var row, col int
		for j := range grid[i] {
			if grid[i][j] != 0 {
				ret++
			}
			row, col = int(math.Max(float64(row), float64(grid[i][j]))), int(math.Max(float64(col), float64(grid[j][i])))
		}
		ret += row + col
	}
	return ret
}
*/
