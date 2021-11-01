package main

/*
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.
給定一個 二進制 的矩陣 A，我們想要水平翻轉他

To flip an image horizontally means that each row of the image is reversed.
翻轉過後的 二進制 陣列A 元素會剛好顛倒

For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
For example, inverting [0, 1, 1] results in [1, 0, 0].

Example 1:
Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]

Example 2:
Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

Notes:
1 <= A.length = A[0].length <= 20
0 <= A[i][j] <= 1
*/

// https://leetcode.com/problems/flipping-an-image/discuss/924616/Go-XOR
func flipAndInvertImage(A [][]int) [][]int {

	// 回傳二維陣列
	result := make([][]int, len(A))
	if len(A) == 0 { // 如果 A 的長度為 0，直接回傳
		return result
	}

	// 假設為方陣，定義row及col的長度
	rows := len(A)
	cols := len(A[0])

	// 賦值給result陣列
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}

	// XOR元素後交換
	for r := 0; r < rows; r++ {
		for c := 0; c <= cols/2; c++ {
			result[r][c], result[r][cols-1-c] = A[r][cols-1-c]^1, A[r][c]^1
		}
	}

	return result
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/flipping-an-image/discuss/184940/Go-easy-solution

func flipAndInvertImage(A [][]int) [][]int {
    for i := 0; i < len(A); i++ {
        row := A[i]
        last := len(row) - 1;
        for j := 0; j <= last / 2; j ++ {
            row[j], row[last - j] = row[last - j] ^ 1, row[j] ^ 1
        }
    }
    return A
}
*/
