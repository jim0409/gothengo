package main

/*
Given a string, you need to reverse the order of characters in each word
within a sentence while still preserving whitespace and initial word order.


Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"


Note:
In the string, each word is separated by single space
and there will not be any extra space in the string.

*/

// https://leetcode.com/problems/reverse-words-in-a-string-iii/discuss/247155/simple-go-solution
func reverseWords(s string) string {
	ls := len(s)
	if ls < 2 {
		return s
	}

	b := []byte(s)
	l, r := 0, 0
	for i, v := range b {

		// 判斷是否具有 space ，有的話就當作一的獨立單字
		// if v is space or come to the end
		if v == ' ' || i == ls-1 {
			r = i - 1 // set the right index
			// but if we get the end of s, we should plus 1 to r, because we subtracted r before
			if i == ls-1 {
				r++
			}

			// 將獨立單字反轉
			// loop between l and r, swap the vlalue on the two indexes
			for l < r {
				b[l], b[r] = b[r], b[l]

				r--
				l++
			}

			// 跳過目前單字長度，因為`空白`跟`單字長度`都是應該被忽略的
			l = i + 1 // set l to i plus 1. because s[i] is space, so we start at the next one
		}
	}
	return string(b)
}

func main() {

}

/*
solution
- https://leetcode.com/problems/reverse-words-in-a-string-iii/discuss/553775/Go%3A-0ms-solution

func reverseWords(s string) string {
    i, j := 0, 0
    res := make([]byte, len(s))

    for j < len(s) {
        if s[j] == ' ' {
            if i == j {
                res[j] = s[j]
            } else {
                reverseCopy(s, res, i, j-1)
                res[j] = s[j]
            }
            i, j = j+1, j+1
            continue
        }
        j++
    }

    if i < j {
        reverseCopy(s, res, i, j-1)
    }
    return string(res)
}

func reverseCopy(s string, res []byte, start, end int) {
    for start < end {
        res[start], res[end] = s[end], s[start]
        start, end = start+1, end-1
    }
    if start == end {
        res[start] = s[start]
    }
}
*/
