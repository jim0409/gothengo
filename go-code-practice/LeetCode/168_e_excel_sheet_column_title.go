package main

/*
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
	...

Example 1:
Input: 1
Output: "A"

Example 2:
Input: 28
Output: "AB"

Example 3:
Input: 701
Output: "ZY"
*/

func convertToTitle(n int) string {
	// A --> Z : 1 --> 26
	// 28 /26 = 1 ... 2 => AB
	// 701 /26 = 26 ... 25 => ZY
	columnTitle := ""
	for n > 0 {
		v := n % 26
		n = n / 26
		if v == 0 {
			v = 26
			n -= 1
		}
		columnTitle = string('A'+byte(v-1)) + columnTitle
	}
	return columnTitle
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/excel-sheet-column-title/discuss/285848/Go-double-100


func convertToTitle(n int) string {
    // A --> Z : 1 --> 26
    // 28 /26 = 1 ... 2 => AB
    // 701 /26 = 26 ... 25 => ZY
    columnTitle := ""
    for n > 0 {
        v := n % 26
        n = n / 26
        if v == 0 {
            v = 26
            n -= 1
        }
        columnTitle = string('A' + byte(v-1)) + columnTitle
    }
    return columnTitle
}
*/
