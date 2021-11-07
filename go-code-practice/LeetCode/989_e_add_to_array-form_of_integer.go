package main

import "strconv"

/*
The array-form of an integer num is an array representing its digits in left to right order.

- For example,
	for `num = 1321`, the array form is [1,3,2,1].

Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.


Example 1:
Input: num = [1,2,0,0], k = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234

Example 2:
Input: num = [2,7,4], k = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455

Example 3:
Input: num = [2,1,5], k = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021

Example 4:
Input: num = [9,9,9,9,9,9,9,9,9,9], k = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000


Constraints:
1 <= num.length <= 10^4
0 <= num[i] <= 9
num does not contain any leading zeros except for the zero itself.
1 <= k <= 10^4
*/

// https://leetcode.com/problems/add-to-array-form-of-integer/discuss/855533/Go-Golang-28-ms-faster-than-92.11
func addToArrayForm(A []int, K int) []int {
	// res 取 A 長度
	res := make([]int, len(A)+1)
	// KLen 將 K 轉成字串，在算出字串長度
	kLen := len(strconv.Itoa(K))
	if kLen > len(A) {
		res = make([]int, kLen+1)
	}

	// carry: 進位
	for carry, i, k := 0, len(A)-1, len(res)-1; k >= 0; carry, i, k = carry/10, i-1, k-1 {
		if i >= 0 {
			carry += A[i]
		}
		// 當 K > 0，對該位數取餘做 10 進位
		if K > 0 {
			carry += K % 10
			// 每次迭代，K取餘數
			K /= 10
		}
		res[k] = carry % 10
	}

	if res[0] == 0 {
		return res[1:]
	}
	return res
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/add-to-array-form-of-integer/discuss/570473/Golang-100-fast

func addToArrayForm(A []int, K int) []int {
	carry := K
	N := 10001
	res := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		tmp := carry
		if 0 <= len(A)-1-(N-1-i) {
			tmp += A[len(A)-1-(N-1-i)]
		}
		if tmp == 0 && 0 > len(A)-1-(N-1-i) {
			if i == N-1 {
				return []int{0}
			}
			return res[i+1:]
		}
		res[i] = tmp % 10
		carry = tmp / 10
	}
	return res
}
*/
