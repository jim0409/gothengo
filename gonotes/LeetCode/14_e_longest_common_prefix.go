package main

import "fmt"

// func longestCommonPrefix(strs []string) string {
// 	// 1. 先找尋最短的字串
// 	// 2. 找尋最短得字串對應的byte是否也存在於其他string
// 	// 3. 依次，比較
// 	short := findshortestString(&strs)
// 	for i := 0; i < len(strs); i++ {
// 		if strs[i] != short {
// 			tmpPattern := short[i : len(short)-1]
// 			comparePatternInString(strs[i], tmpPattern)
// 		}
// 	}

// 	return short
// }

// // find shortest string 會找尋最短的字串，以這個字串來尋找接下來的字串陣列中是否也有符合的
// func findshortestString(ptrstrs *[]string) string {
// 	if len(*ptrstrs) == 0 {
// 		return ""
// 	}

// 	res := (*ptrstrs)[0]
// 	for _, s := range *ptrstrs {
// 		if len(res) > len(s) {
// 			res = s
// 		}
// 	}
// 	return res
// }

// func comparePatternInString(strs string, sub string) bool {
// 	exist := false
// 	lenSub := len(sub)
// 	/*
// 		考慮比較
// 		pattern: ab
// 		字串: adabde
// 		順序: ad -> da -> ab
// 		(由左到右，依序比較)
// 	*/
// 	for i := 0; i < len(strs)-lenSub-1; i++ {
// 		if strs[i:i+lenSub] == sub {
// 			exist = true
// 		}
// 	}

// 	return exist
// }

/*
	Write a function to find the longest common prefix string
	amongst an array of strings.

	If there is no common prefix, return an empty string ""
*/

func main() {
	// Input ["flower", "flow", "flight"]
	// Output "fl"

	// s := "testing"
	// fmt.Println(comparePatternInString(s, "te"))

	arrStr := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(arrStr))

	// Input ["dog", "racecar", "car"]
	// Output ""
	// Explanation: there is no common prefix among the input strings.

}

// 參考解答 ... 一題easy...解起來跟medium有得拼... 很考驗golang slice的一題，根本就ZzZz...難怪負讚那麼多... 漏看一個prefix實作差很多啊...
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 預存一個 pre 的 byte矩陣
	pre := []byte{}

	// 拿第一個字串當作pattern
	for j := 0; j < len(strs[0]); j++ {

		// 第二個字串開始做比對
		for i := 1; i < len(strs); i++ {

			// 如果 字串長度為0(或者小於比對字串的長度) 或者 第一個字串的 pattern[0][j] 不等於 字串 [i][j] 回傳 string(pre)
			// j >= len(strs[i]) : 基本不用比，因為除非字串為""，否則都會為false => 考驗，如果比對的pattern 大於等於要比對的字串，那麼就不用再往下比了
			// strs[0][j] != strs[i][j] => 這個地方，i是會遞增的，其實是在比心酸的 ... 因為有prefix的邀前，所以前綴一定要相同。也就是 不會出現情況為 `xxas` vs `axxas` (prefix不同)
			if j >= len(strs[i]) || strs[0][j] != strs[i][j] {
				return string(pre)
			}
		}
		pre = append(pre, strs[0][j])
	}
	return string(pre)
}
