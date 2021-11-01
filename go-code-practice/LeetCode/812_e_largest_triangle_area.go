package main

import "math"

/*
You have a list of points in the plane.
Return the area of the largest triangle that can be formed by any 3 of the points.

Example:
Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2
Explanation:
The five points are show in the figure below.
The red triangle is the largest.

.
.
. . .

Notes:
- 3 <= points.length <= 50.
- No points will be duplicated.
- -50 <= points[i][j] <= 50.
- Answers within 10^-6 of the true value will be accepted as correct.
*/

// https://leetcode.com/problems/largest-triangle-area/discuss/290668/Golang-solution-beat-100-people.
func largestTriangleArea(points [][]int) float64 {
	res := 0.0 // 緩存最大值
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				x, y, z := points[i], points[j], points[k]
				x1 := x[0]
				y1 := x[1]
				x2 := y[0]
				y2 := y[1]
				x3 := z[0]
				y3 := z[1]
				newArea := math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2) / 2.0)
				res = math.Max(res, newArea)
			}
		}
	}
	return res
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/largest-triangle-area/discuss/259655/simple-golang-solution

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	var ans float64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				ans = max(ans, area(points[i], points[j], points[k]))
			}
		}

	}
	return ans
}

func area(p []int, q []int, r []int) float64 {
	y := p[0]*q[1] + q[0]*r[1] + r[0]*p[1] - (p[1]*q[0] + q[1]*r[0] + r[1]*p[0])

	return 0.5 * float64(abs(y))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
*/
