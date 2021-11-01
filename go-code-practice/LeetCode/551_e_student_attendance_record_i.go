package main

/*
You are given a string representing an attendance record for a student.
The record only contains the following three characters:

'A' : Absent.
'L' : Late.
'P' : Present.

A student could be rewarded
if his attendance record doesn't contain more than one 'A' (absent)
or more than two continuous 'L' (late).

You need to return whether the student could be rewarded
according to his attendance record.

Example 1:
Input: "PPALLP"
Output: True

Example 2:
Input: "PPALLL"
Output: False

*/

func checkRecord(s string) bool {
	for A, L, i := 0, 0, 0; i < len(s); i++ {
		if s[i] == 'A' {
			A++
		}
		if s[i] == 'L' {
			L++
		} else {
			L = 0
		}
		if A > 1 || L > 2 {
			return false
		}
	}
	return true
}

func main() {

}

/*
solution
- https://leetcode.com/problems/student-attendance-record-i/discuss/101602/Go-0ms

func checkRecord(s string) bool {
	return !(strings.Count(s, `A`) > 1) && !strings.Contains(s, `LLL`)
}
*/
