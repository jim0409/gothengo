package main

/*
You are given a map in form of a two-dimensional integer grid
where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally)
The grid is completely surrounded by water, and there is exactly one island
(i.e., one or more connected land cells).


The island doesn't have "lakes"
(water inside that isn't connected to the water around the island).
One cell is a square with side length 1.

The grid is rectangular, width and height don't exceed 100.
Determine the perimeter of the island.


Example:

Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16
*/

// 算陸地`1`面積所繞成的周長
func islandPerimeter(grid [][]int) int {

	cn := 0

	// line 計算二維陣列中的 陣列總數
	for line := 0; line < len(grid); line++ {
		// row 計算 每一個二維陣列中的`第一個值`
		for row := 0; row < len(grid[0]); row++ {
			// 如果 grid[line][row]為0，表示該位置是水源，直接跳過
			if grid[line][row] == 0 {
				continue
			}

			UP := !(line == 0)             //false
			DOWN := !(line == len(grid)-1) //false

			LEFT := !(row == 0)               //false
			RIGHT := !(row == len(grid[0])-1) //false

			tmp := 4

			if UP {
				tmp = tmp - grid[line-1][row]
			}
			if DOWN {
				tmp = tmp - grid[line+1][row]
			}
			if LEFT {
				tmp = tmp - grid[line][row-1]
			}
			if RIGHT {
				tmp = tmp - grid[line][row+1]
			}
			cn = cn + tmp
		}
	}
	return cn
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/island-perimeter/discuss/527503/Go-O(mn)-solution


func islandPerimeter(grid [][]int) int {
	var lands, neighbours int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			lands++
			if i != len(grid)-1 && grid[i+1][j] == 1 {
				neighbours++
			}
			if j != len(grid[0])-1 && grid[i][j+1] == 1 {
				neighbours++
			}
		}
	}
	return 4*lands - 2*neighbours
}
*/

/*
solution2:
- https://leetcode.com/problems/island-perimeter/discuss/95061/Go-dfs-solution


func islandPerimeter(grid [][]int) int {
    if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    r, c := len(grid), len(grid[0])
    v := make([][]bool, r)
    for i := 0; i < r; i++ {
        v[i] = make([]bool, c)
    }
    p := 0
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            p += dfs(grid, v, r, c, i, j)
        }
    }
    return p
}

func dfs(grid [][]int, v [][]bool, r, c, i, j int) int {
    if i < 0 || i >= r || j < 0 || j >= c || grid[i][j] == 0 || v[i][j] {
        return 0
    }
    v[i][j] = true
    p := 0
    if i == 0 || grid[i - 1][j] == 0 {
        p++
    }
    if j == 0 || grid[i][j - 1] == 0 {
        p++
    }
    if i == r - 1 || grid[i + 1][j] == 0 {
        p++
    }
    if j == c - 1 || grid[i][j + 1] == 0 {
        p++
    }
    return p + dfs(grid, v, r, c, i - 1, j) + dfs(grid, v, r, c, i + 1, j) +
            dfs(grid, v, r, c, i, j - 1) + dfs(grid, v, r, c, i, j + 1)
}
*/

/*
solution3:
- https://leetcode.com/problems/island-perimeter/discuss/226121/100-use-go


func islandPerimeter(grid [][]int) int {

    cn := 0
    for line:=0;line<len(grid);line++ {
        for row:=0;row<len(grid[0]);row++ {
            if grid[line][row] == 0 {
                continue
            }

            UP    := !(line == 0 ) //false
            DOWN  := !(line == len(grid) - 1) //false

            LEFT  := !(row == 0 ) //false
            RIGHT := !(row == len(grid[0]) - 1) //false

            tmp := 4

            if UP {
                tmp = tmp - grid[line -1][row]
            }
            if DOWN {
                tmp = tmp - grid[line +1][row]
            }
            if LEFT {
                tmp = tmp - grid[line][row -1]
            }
            if RIGHT {
                tmp = tmp - grid[line][row +1]
            }
            cn = cn + tmp
        }
    }
    return cn
}
*/
