package main

import "fmt"

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil

	} else if numRows == 1 {
		return [][]int{{1}}

	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}

	}

	triange := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		// 定義 回傳的 ps 三角形
		triange = append(triange, make([]int, i+1))

		// 先宣告 ps 第一跟最後的數字皆為 1
		triange[i][0], triange[i][len(triange[i])-1] = 1, 1

		for j := 1; j < len(triange[i])-1; j++ {
			// 三角形最底層的數字為前一層的對應位置總和
			triange[i][j] = triange[i-1][j-1] + triange[i-1][j]
		}
	}

	return triange
}

func main() {
	n := 4
	fmt.Println(generate(n))
}

/*
solution
https://leetcode.com/problems/pascals-triangle/discuss/302026/Go-simple-solution

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	} else if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	triange := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		triange = append(triange, make([]int, i+1))
		triange[i][0], triange[i][len(triange[i])-1] = 1, 1
		for j := 1; j < len(triange[i])-1; j++ {
			triange[i][j] = triange[i-1][j-1] + triange[i-1][j]
		}
	}
	return triange
}
*/
