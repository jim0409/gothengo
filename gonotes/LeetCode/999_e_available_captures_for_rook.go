package main

/*
On an 8 x 8 chessboard,

there is exactly one white rook 'R' and some number of white bishops 'B', black pawns 'p', and empty squares '.'.

When the rook moves, it chooses one of four cardinal directions (north, east, south, or west),

then moves in that direction until it chooses to stop, reaches the edge of the board,

captures a black pawn, or is blocked by a white bishop.

A rook is considered attacking a pawn if the rook can capture the pawn on the rook's turn.

The number of available captures for the white rook is the number of pawns that the rook is attacking.

Return the number of available captures for the white rook.

Example 1:
Input: board = [[".",".",".",".",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".","R",".",".",".","p"],
		[".",".",".",".",".",".",".","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".",".",".",".",".","."]]
Output: 3
Explanation: In this example, the rook is attacking all the pawns.


Example 2:
Input: board = [[".",".",".",".",".",".",".","."],
		[".","p","p","p","p","p",".","."],
		[".","p","p","B","p","p",".","."],
		[".","p","B","R","B","p",".","."],
		[".","p","p","B","p","p",".","."],
		[".","p","p","p","p","p",".","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".",".",".",".",".","."]]
Output: 0
Explanation: The bishops are blocking the rook from attacking any of the pawns.


Example 3:
Input: board = [[".",".",".",".",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		["p","p",".","R",".","p","B","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".","B",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".",".",".",".",".","."]]
Output: 3
Explanation: The rook is attacking the pawns at positions b5, d6, and f5.

Constraints:
- board.length == 8
- board[i].length == 8
- board[i][j] is either 'R', '.', 'B', or 'p'
- There is exactly one cell with board[i][j] == 'R'
*/

// https://leetcode.com/problems/available-captures-for-rook/discuss/243464/Golang-solution-beats-100(Can't-get-simpler-than-this!!!)
/*
找到 Rook 在棋盤中的位置，根據 Rook 的 4 個方向延伸，做情境判斷
判斷，如果遇到 `p` 就加一後跳出迴圈，遇到 `B` 則不直接跳出迴圈
*/
func numRookCaptures(board [][]byte) int {
	count := 0
	i := 0
	j := 0
	for a, row := range board {
		for b, cell := range row {
			if string(cell) == "R" {
				i = a
				j = b
				break
			}
		}
	}
	for _, directions := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		x := i + directions[0]
		y := j + directions[1]
		for x >= 0 && y >= 0 && x < len(board) && y < len(board[0]) && (board[x][y] == '.' || board[x][y] == 'p') {
			if board[x][y] == 'p' {
				count += 1
				break
			}
			x += directions[0]
			y += directions[1]
		}
	}
	return count
}

func main() {

}
