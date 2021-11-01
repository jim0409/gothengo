package main

/*
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
	...

Example 1:
Input: "A"
Output: 1

Example 2:
Input: "AB"
Output: 28

Example 3:
Input: "ZY"
Output: 701
*/

func titleToNumber(s string) int {
	col := 0
	for _, r := range s {
		cur := int(r - 'A' + 1)
		col = col*26 + cur
	}
	return col
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/excel-sheet-column-number/discuss/182719/Go-Solution

func titleToNumber(s string) int {
    col := 0
    for _,r := range s {
        cur := int(r-'A'+1)
        col = col*26+cur
    }
    return col
}
*/
