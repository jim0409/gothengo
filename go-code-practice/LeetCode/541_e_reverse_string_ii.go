package main

/*
Given a string and an integer k,
you need to reverse the first k characters for every 2k characters counting
from the start of the string.

如果字串的長度小於k，那麼將所有字串反轉過來
If there are less than k characters left, reverse all of them.

如果字串中的字元的介於k到2k之間，那麼將前k個元素互換。而之後的數則不變
If there are less than 2k but greater than or equal to k characters,
then reverse the first k characters and left the other as original.


Example:
Input: s = "abcdefg", k = 2
			12
			  12
			  	123
			=> 2, 4, 6, 8

Output: "bacdfeg"
=> 12 互換，34互換

Restrictions:
The string consists of lower English letters only.
Length of the given string and k will in the range [1, 10000]
*/

// https://leetcode.com/problems/reverse-string-ii/discuss/509983/Go-simple-100-solution
func reverseStr(s string, k int) string {
	res := []rune(s) // 回傳字串
	h := 0           // (初始)計算頭
	t := k - 1       // (初始)計算尾
	n := 0           // (乘以k的)倍數迭代值

	for n*k < len(s) {
		if t > len(s)-1 {
			t = len(s) - 1
		}

		for h < t { // 如果 頭小於尾 則互換
			res[h], res[t] = res[t], res[h]
			h++
			t--
		}

		n += 2          // n 每次 加 2
		h = n * k       // 下一次的計算頭為 n*k
		t = (n+1)*k - 1 // 下一次的計算尾為 (n+1)*k -1
	}

	return string(res)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/reverse-string-ii/discuss/418106/Go-0ms-double-100

func reverseStr(s string, k int) string {
    b := []byte(s)
    total, remaining := len(s), len(s)
    for remaining > 0 {
        flag, distance := 0, 0
        if remaining >= k {
            flag = k/2
            distance = k-1
        } else {
            flag = remaining/2
            distance = remaining-1
        }
        for i := 0; i < flag; i++ {
            b[total-remaining+i], b[total-remaining+distance-i] = b[total-remaining+distance-i], b[total-remaining+i]
        }
        remaining -= 2*k
    }
    return string(b)
}
*/
