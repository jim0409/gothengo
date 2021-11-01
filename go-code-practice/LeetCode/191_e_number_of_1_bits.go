package main

/*
Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).



Example 1:

Input: 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
Example 2:

Input: 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.
Example 3:

Input: 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.


Note:

Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3 above the input represents the signed integer -3.


Follow up:

If this function is called many times, how would you optimize it?
*/

// 寫法沒利用到 golang operator 而且要額外宣告一個 string
// func hammingWeight(num uint32) int {
// 	res := 0
// 	tmp := fmt.Sprintf("%b", num)
// 	for i := 0; i < len(tmp); i++ {
// 		if tmp[i] == '1' {
// 			res++
// 		}
// 	}
// 	return res
// }

func hammingWeight(num uint32) int {
	var result int

	// 從 32位元開始算~
	for i := 31; i >= 0; i-- {
		// result 為 num 的每個位元判斷後++
		// uint(i)，表示 num 要往右邊移幾個位元
		// (num ...) & 1，如果右移後該位元為1就返回true，再把true轉型為1加回result
		result += int((num >> uint(i)) & 1)
	}

	return result
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/number-of-1-bits/discuss/244614/go-solution

func hammingWeight(num uint32) int {
	var result int

	for i := 31; i >= 0; i-- {
		result += int((num >> uint(i)) & 1)
	}

	return result

}
*/
