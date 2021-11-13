package main

/*
Implement function ToLowerCase() that has a string parameter str,
and returns the same string in lowercase.


Example 1:
Input: "Hello"
Output: "hello"

Example 2:
Input: "here"
Output: "here"

Example 3:
Input: "LOVELY"
Output: "lovely"
*/

func toLowerCase(str string) string {
	runes := []rune(str)
	diff := 'a' - 'A'
	for i := 0; i < len(str); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += diff
		}
	}
	return string(runes)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/to-lower-case/discuss/155064/Go-code

func toLowerCase(str string) string {
    r:=[]rune(str)
    for i:=0;i<len(r);i++{
        if r[i]<=90&&r[i]>=65 {
            r[i]+=32
        }
    }
    return string(r)
}
*/
