package main

import (
	"fmt"
)

/*
Reverse bits of a given 32 bits unsigned integer.


Example 1:

Input: 00000010100101000001111010011100
Output: 00111001011110000010100101000000
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
Example 2:

Input: 11111111111111111111111111111101
Output: 10111111111111111111111111111111
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.


Note:

Note that in some languages such as Java, there is no unsigned integer type. In this case, both input and output will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above the input represents the signed integer -3 and the output represents the signed integer -1073741825.
*/

// num 為一個 uint32 的數字 e.g. 123 = 0...01111011
func reverseBits(num uint32) uint32 {
	// 宣告一個 uint32 的 result
	result := uint32(0)
	// log.Printf("%b\n", result) // 0

	// 從nums的第一個數字開始做，做到第32個位元
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		// log.Printf("result is %b\n", result)

		num = num >> 1
		// log.Printf("num is %b\n", num)
	}
	return result
}

func main() {
	a := uint32(123)
	// fmt.Println(reverseBits(a))
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", reverseBits(a))
}

/*
solution:
- https://leetcode.com/problems/reverse-bits/discuss/293390/Go-solution-0ms

func reverseBits(num uint32) uint32 {
    result := uint32(0)
    for i:=0; i<32; i++ {
        result = (result<<1)|(num&1)
        num = num>>1
    }
    return result
}
*/
