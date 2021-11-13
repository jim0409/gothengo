package main

import "fmt"

/*
You are given an array of n strings strs, all of the same length.

The strings can be arranged such that there is one on each line, making a grid.
For example, strs = ["abc", "bce", "cae"] can be arranged as:

abc
bce
cae

You want to delete the columns that are not sorted lexicographically.
In the above example (0-indexed), columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted
while column 1 ('b', 'c', 'a') is not, so you would delete column 1.

Return the number of columns that you will delete.


Example 1:
Input: strs = ["cba","daf","ghi"]
Output: 1
Explanation: The grid looks as follows:
  cba
  daf
  ghi
Columns 0 and 2 are sorted, but column 1 is not, so you only need to delete 1 column.

Example 2:
Input: strs = ["a","b"]
Output: 0
Explanation: The grid looks as follows:
  a
  b
Column 0 is the only column and is sorted, so you will not delete any columns.

Example 3:
Input: strs = ["zyx","wvu","tsr"]
Output: 3
Explanation: The grid looks as follows:
  zyx
  wvu
  tsr
All 3 columns are not sorted, so you will delete all 3.


Constraints:
n == strs.length
1 <= n <= 100
1 <= strs[i].length <= 1000
strs[i] consists of lowercase English letters.
*/

/*
題意:
給定一個字串陣列並且觀察欄字元，觀察該欄字元是否為有序
假設該欄非有序的，則回傳該欄的編號
註: 假設每個字的長度均相同
e.g.
[ "cba", "daf", "ghi" ]
c > d > g => 第 0 欄為有序的
d > a > h => 第 1 欄為無序的
a > f > i => 第 2 欄為有序的
*/

func minDeletionSize(A []string) int {
	var ret int
	for j := 0; j < len(A[0]); j++ {
		for i := 1; i < len(A); i++ {
			if A[i][j] < A[i-1][j] {
				ret++
				break
			}
		}
	}
	return ret
}

func main() {
	var t1 = []string{"zyx", "wvu", "tsr"}
	fmt.Println(minDeletionSize(t1))
}

/*
refer:
- https://leetcode.com/problems/delete-columns-to-make-sorted/discuss/739363/Go.-A-crappy-solution-for-a-crappy-problem.

func minDeletionSize(A []string) int {
    res := 0
    for col := 0; col < len(A[0]); col++ {
        for row := 1; row < len(A); row++ {
            if A[row][col] < A[row-1][col] {
                res++
                break
            }
        }
    }
    return res
}

*/
