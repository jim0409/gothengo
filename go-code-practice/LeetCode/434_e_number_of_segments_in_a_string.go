package main

import "fmt"

/*
Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.
計算一個字串中的`分片詞`，那些連續的詞(以空白間隔)

Please note that the string does not contain any non-printable characters.
假設字串中不存在不可打印的字元

Example:
Input: "Hello, my name is John"
Output: 5
*/

// func countSegments(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	count := 0
// 	if s[0] != ' ' {
// 		count++
// 	}
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == ' ' && s[i-1] != ' ' {
// 			count++
// 		}

// 		if s[i] == ' ' && s[i-1] == ' ' {
// 			break
// 		}
// 	}

// 	if s[len(s)-1] == ' ' {
// 		count--
// 	}

// 	return count
// }

/*
1. 先循遍所有的空白，並把空白的個數記錄下來。當作真實開始計算時的起點
2. 由該起點開始往後算
3. 如該起點往後算遇到的非空白字元，將位移起點i++
4. 將"所有空白數"加上"第r次迭代的奇點後非空白字元"，作為下次起點開始迭代運算
5. 迭代到起點大於字串總長度

Q: 為何要計算便歷空白個數

Q: 為何可以用位移後的起點來計算

Q: 為何要重複運算

*/
func countSegments(s string) int {
	i, n, r := 0, len(s), 0
	// 先遍歷一遍找到所有空白的個數
	for i < n && s[i] == ' ' {
		i++
	}
	// 從 i 開始往後算
	for i < n {
		r++

		// 計算所有非空白字元的數量，並位移起點i
		for i < n && s[i] != ' ' {
			i++
		}

		// 在計算完所有非空白字元的數量後，再次計算空白數量。且位移起點i
		for i < n && s[i] == ' ' {
			i++
		}
	}
	return r
}

func main() {
	s1 := "Hello, my name is John"
	fmt.Println(countSegments(s1))

	s2 := "                "
	fmt.Println(countSegments(s2))

	s3 := "Of all the gin joints in all the towns in all the world,   "
	fmt.Println(countSegments(s3))
}

/*
solution:
- https://leetcode.com/problems/number-of-segments-in-a-string/discuss/133320/Simple-Go-solution


func countSegments(s string) (total int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	for i:=1 ; i< len(s) ; i++ {
		if s[i] ==' ' && s[i-1] != ' ' {
			total++
		}
	}
	return total+1
}
*/

/*
solution2:
- https://leetcode.com/problems/number-of-segments-in-a-string/discuss/483679/go-brute-force


// 0ms 100.00% 1.9MB 100.00%
// brute force
// O(N) O(1)
func countSegments(s string) int {
  i, n, r := 0, len(s), 0
  // skip white spaces
  for i < n && s[i] == ' ' {
    i++
  }
  for i < n {
    r++
    for i < n && s[i] != ' ' {
      i++
    }
    for i < n && s[i] == ' ' {
      i++
    }
  }
  return r
}
*/

/*
solution3:
- https://leetcode.com/problems/number-of-segments-in-a-string/discuss/412170/Go-golang-0ms-clean-solution

func countSegments(s string) int {

    if len(s) == 0 { return 0 }
	count := 1

	for i := 0; i < len(s); i++ {
        if i == 0 && s[i] == ' ' {
            count = 0
        }
		if i < len(s)-1 && s[i] == ' ' && s[i+1] != ' ' {
			count++
		}
	}
	return count
}
*/
