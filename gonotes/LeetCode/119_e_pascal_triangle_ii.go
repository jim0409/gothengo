package main

import "fmt"

/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?
*/

func getRow(rowIndex int) []int {
	res := generate(rowIndex + 1)
	return res[len(res)-1]
}

func main() {
	n := 4
	// fmt.Println(generate(n)
	fmt.Println(getRow(n))
}

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

/*
solution
https://leetcode.com/problems/pascals-triangle-ii/discuss/539696/Go%3A-Recursive-Solution-100

func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    } else if rowIndex == 1 {
        return []int{1, 1}
    }
    return recursive([]int{1, 1}, 2, rowIndex)
}

func recursive(prev []int, current, target int) []int {
    row := make([]int, current+1)
    row[0], row[len(row)-1] = 1, 1

    for i := 1; i < current; i++ {
        row[i] = prev[i-1] + prev[i]
    }

    if current == target {
        return row
    } else {
        return recursive(row, current+1, target)
    }
}
*/
