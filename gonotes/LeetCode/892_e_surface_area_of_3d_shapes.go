package main

import "math"

/*
You are given an `n x n` grid where you have placed some `1 x 1 x 1` cubes.
Each value v = grid[i][j] represents a tower of v cubes placed on top of cell (i, j).

After placing these cubes, you have decided to glue any directly adjacent cubes to each other,
forming several irregular 3D shapes.

Return the total surface area of the resulting shapes.

Note: The bottom face of each shape counts toward its surface area.


Example 1:
Input: grid = [[2]]
Output: 10

Example 2:
Input: grid = [[1,2],[3,4]]
Output: 34

Example 3:
Input: grid = [[1,0],[0,2]]
Output: 16

Example 4:
Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
Output: 32

Example 5:
Input: grid = [[2,2,2],[2,1,2],[2,2,2]]
Output: 46


Constraints:

n == grid.length
n == grid[i].length
1 <= n <= 50
0 <= grid[i][j] <= 50
*/

// refer:
// - https://leetcode.com/problems/surface-area-of-3d-shapes/discuss/365454/Go-golang-0ms-clean-solution
// 算出物體表面面積
func surfaceArea(grid [][]int) int {

	tmp := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				tmp += 2
			}
			if i == 0 {
				tmp += grid[i][j]
			}
			if j == 0 {
				tmp += grid[i][j]
			}
			if i == len(grid)-1 {
				tmp += grid[i][j]
			}
			if j == len(grid[i])-1 {
				tmp += grid[i][j]
			}
			if i < len(grid)-1 {
				tmp += int(math.Abs(float64(grid[i][j] - grid[i+1][j])))
			}
			if j < len(grid[i])-1 {
				tmp += int(math.Abs(float64(grid[i][j] - grid[i][j+1])))
			}
		}
	}
	return tmp
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/surface-area-of-3d-shapes/discuss/513467/Go-double-100-solution

func surfaceArea(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			area += grid[i][j]*4 + 2
			if i > 0 {
				area -= min(grid[i-1][j], grid[i][j]) * 2
			}
			if j > 0 {
				area -= min(grid[i][j-1], grid[i][j]) * 2
			}
		}
	}
	return area
}

func min(values ...int) int {
	minValue := math.MaxInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

*/
