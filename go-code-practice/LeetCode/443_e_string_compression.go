package main

import "strconv"

/*
Given an array of characters, compress it in-place.
The length after compression must always be smaller than or equal to the original array.
Every element of the array should be a character (not int) of length 1.
After you are done modifying the input array in-place, return the new length of the array.
給定一個字元陣列，嘗試把它壓縮後再打印出來


Follow up:
Could you solve it using only O(1) extra space?


Example 1:

Input:
["a","a","b","b","c","c","c"]

Output:
Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]

Explanation:
"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".
"aa"可以取代為"a2"(因為"a"出現兩次)
*/

func compress(chars []byte) int {
	// ret 為最後要回傳的字串長度
	var ret int
	// i,j 分別為字串的`起點`跟`終點`
	for i, j := 0, 0; j < len(chars); i = j {
		// 算出字元陣列中連續出現字元長度的最大值
		// j 的最大值 (b,a,a,b) => j = 0
		// j 的最大值 (a,a,b,b) => j = 2
		for ; j < len(chars) && chars[i] == chars[j]; j++ {
		}

		chars[ret] = chars[i]
		ret++
		if j-i > 1 {
			// cnt 為 `兩個相同的字元最長的距離`
			cnt := strconv.Itoa(j - i)
			// 考慮最長的距離超過兩位數 ["a", "1", "2"]
			for k := 0; k < len(cnt); k++ {
				chars[ret] = cnt[k]
				ret++
			}
		}
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/string-compression/discuss/409939/Go-golang-clean-solution


func compress(chars []byte) int {

    scan := 0
    write := 0
    l := len(chars)

    for scan < l {
        count := 0
        chars[write] = chars[scan]
        for scan < l && chars[write] == chars[scan] {
            count++
            scan++
        }
        if count > 1 {
            tmp := fmt.Sprintf("%d", count)
            for _, c := range []byte(tmp) {
                write++
                chars[write] = c
            }
        }
        write++
    }
    return write
}
*/
