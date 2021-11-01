package main

/*
Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother
to make the gray scale of each cell becomes the average gray scale
(rounding down) of all the 8 surrounding cells and itself.
If a cell has less than 8 surrounding cells, then use as many as you can.


Example 1:
Input:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
Output:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
Explanation:
For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
For the point (1,1): floor(8/9) = floor(0.88888889) = 0


Note:
The value in the given matrix is in the range of [0, 255].
The length and width of the given matrix are in the range of [1, 150].
*/

func imageSmoother(M [][]int) [][]int {
	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	ret := make([][]int, len(M))
	for i := range M {
		ret[i] = make([]int, len(M[i]))
		for j := range M[i] {
			sum, cnt := M[i][j], 1
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if 0 <= x && x <= len(M)-1 && 0 <= y && y <= len(M[i])-1 {
					sum, cnt = sum+M[x][y], cnt+1
				}
			}
			ret[i][j] = sum / cnt
		}
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/image-smoother/discuss/507195/Go-44ms-100-solution

func imageSmoother(M [][]int) [][]int {
    processedImage := make([][]int, len(M))
    for i := range processedImage {
        processedImage[i] = make([]int, len(M[0]))
    }
    for i := range processedImage {
        for j := range processedImage[i] {
            processedImage[i][j] = average(M, i, j)
        }
    }
    return processedImage
}

func average(M [][]int, i, j int) int {
    count, summary := 0, 0
    for m := i-1; m <= i+1; m++ {
        for n := j-1; n <= j+1; n++ {
            if v, exist := gray(M, m, n); exist {
                count++
                summary += v
            }
        }
    }
    return summary / count
}

func gray(M [][]int, i, j int) (int, bool) {
    if i < 0 || i >= len(M) || j < 0 || j >= len(M[0]) {
        return 0, false
    }
    return M[i][j], true
}
*/
