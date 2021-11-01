package main

import "fmt"

/*
Count the number of prime numbers less than a non-negative number, n.

Example:

Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

// https://medium.com/coding-by-dumbbell/204-count-primes-849c59c9a503
// 解釋為什麼要這樣計算
func countPrimes(n int) int {
	// 如果 n 小於 2 表示 n 是 質數，直接返回
	if n < 2 {
		return 0
	}

	// 宣告一個儲存 prime 數的布林陣列
	dict := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !dict[i] {
			count++

			// 每一個 質數 對應的倍率 都應該被刪去
			// 因為 質數=質數*1
			for j := 2; i*j < n; j++ {
				dict[i*j] = true
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(19))
}

/*
solution:
- https://leetcode.com/problems/count-primes/discuss/122509/12-ms-Go-solution


func countPrimes(n int) int {
    if n < 2 {
        return 0
    }
    dict := make([]bool, n)
    count := 0
    for i:=2; i<n; i++ {
        if !dict[i] {
            count++
            for j:=2; i*j < n; j++ {
                dict[i*j] = true
            }
        }
    }
    return count
}
*/
