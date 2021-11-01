package main

import "fmt"

/*
We are to write the letters of a given string S, from left to right into lines.
假設我們在寫一封信，由左至右

Each line has maximum width 100 units,
每一行有100個單位，
and if writing a letter would cause the width of the line to exceed 100 units, it is written on the next line.
假設信件的字數超過該行的單位大小，則需要轉移至下一行

We are given an array widths,
給定一個"字體"寬度標準的矩陣
an array where widths[0] is the width of 'a', widths[1] is the width of 'b', ..., and widths[25] is the width of 'z'.
矩陣中第0個元素表示'a'，第1個元素表示'b'，以此類推至第25個元素表示'z'

Now answer two questions:
how many lines have at least one character from S,
and what is the width used by the last such line? Return your answer as an integer list of length 2.
假設有一個句話`S`，以及其對應內每個字符所需要的'字體寬度'的矩陣，請回答要花費多少行，以及下一行要使用多少字

Example :
Input:
widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "abcdefghijklmnopqrstuvwxyz"
Output: [3, 60]
Explanation:
All letters have the same length of 10. To write all 26 letters,
we need two full lines and one line with 60 units.

Example :
Input:
widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
S = "bbbcccdddaaa"
Output: [2, 4]
Explanation:
All letters except 'a' have the same length of 10, and
"bbbcccdddaa" will cover 9 * 10 + 2 * 4 = 98 units.
For the last 'a', it is written on the second line because
there is only 2 units left in the first line.
So the answer is 2 lines, plus 4 units in the second line.


Note:
The length of S will be in the range [1, 1000].
S will only contain lowercase letters.
widths is an array of length 26.
widths[i] will be in the range of [2, 10].
*/

// https://leetcode.com/problems/number-of-lines-to-write-string/discuss/352847/Go-golang-0ms-clean-solution
func numberOfLines(widths []int, S string) []int {

	// 先枚舉每個字元所花費的寬度
	h := make(map[rune]int)
	tmp := "abcdefghijklmnopqrstuvwxyz"
	lines := 1
	n := 0

	// 將字元的寬度做儲存
	for i, v := range tmp {
		h[v] = widths[i]
	}

	for _, v := range S {
		n = n + h[v]
		// 當計算的寬度大於100時，增加行數
		if n > 100 {
			lines++
			// 加完行數後，讓n重新回去迭代
			n = h[v]
		}
	}

	return []int{lines, n}

}

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := " abcdefghijklmnopqrstuvwxyz"
	fmt.Sprintf("%v", numberOfLines(widths, s))

}

/*
refer:
- https://leetcode.com/problems/number-of-lines-to-write-string/discuss/159476/Golang-Go-0ms-solution

func numberOfLines(widths []int, s string) []int {
	curLine := 1
	numOnLine := 0
	for _, r := range s {
		curWidth := widths[int(r-'0')-49] // a - '0' is 49
		if numOnLine+curWidth < 101 {
			numOnLine += curWidth
		} else {
			numOnLine = curWidth
			curLine++
		}
	}
	return []int{curLine, numOnLine}
}
*/
