package main

/*
The Hamming distance between two integers is the number of positions at
which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 2^31.

Example:
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
	   ^   ^

The above arrows point to positions where the corresponding bits are different.
*/

func hammingDistance(x int, y int) int {
	hd := x ^ y // x & y 的bitwise，取交集
	r := 0
	for hd != 0 {
		if hd&1 == 1 {
			r++
		}
		hd >>= 1
	}
	return r
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/hamming-distance/discuss/123940/0ms-Go-solution


// x XOR y will set 1 where x and y have different bit
// let z = x XOR y, z - 1 will flip the last set bit in z to 0 and following bits to 1, due to borrowing.
// e.g. 00101000
//    - 00000001
//    = 00100111
// then z & (z-1) will set all the flipped bits to 0, as 1 & 0 = 0
//      00101000
//    & 00100111
//    = 00100000
// therefore we repeat this until z becomes 0, where all set bits are removed.
// return the repeated times as number of set bits

func hammingDistance(x int, y int) int {
    z, d := x ^ y, 0
    for z!=0 {
        z &= z-1
        d++
    }
    return d
}
*/
