package main

/*
Given a positive integer, output its complement number.
The complement strategy is to flip the bits of its binary representation.


Example 1:
Input: 5
Output: 2
Explanation: The binary representation of 5 is 101
(no leading zero bits), and its complement is 010. So you need to output 2.


Example 2:
Input: 1
Output: 0
Explanation: The binary representation of 1 is 1
(no leading zero bits), and its complement is 0. So you need to output 0.


Note:

The given integer is guaranteed to fit within the range of a 32-bit signed integer.
You could assume no leading zero bit in the integer’s binary representation.
This question is the same as 1009: https://leetcode.com/problems/complement-of-base-10-integer/
*/

// 算bitwise, 將原本二進位表示的數字，應該為1的轉為0，反之0轉為1
func findComplement(num int) int {
	temp := 1                      // 0000 0000 0000 0001
	for ; temp <= num; temp *= 2 { // 0000 0000 0000 0010 <-- 近一位來比較
	}
	return (temp - 1) ^ num // 當temp > num 時，temp 應該為 0...0 1 0..0，此時在跟`num`做XOR，算恰好一個為1的位置
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/number-complement/discuss/467826/go-clean-code-0ms-beats-100

func findComplement(num int) int {
	mask := num
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	return num ^ mask
}
*/

/*
solution2:
- https://leetcode.com/problems/number-complement/discuss/225428/100-use-go


func findComplement(num int) int {
    bts := []rune(fmt.Sprintf("%b",num))
    cn := 0
    num_len := len(bts)
    for k:=0;k< num_len;k++ {
        cn = cn + int(bts[k] ^ 1 - '0')*int(math.Pow(2,float64(num_len - k -1)))
    }
    return cn
}
*/
