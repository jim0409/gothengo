package main

/*
We have two special characters.
The first character can be represented by one bit 0.
The second character can be represented by two bits (10 or 11).

Now given a string represented by several bits.
Return whether the last character must be a one-bit character or not.
The given string will always end with a zero.


Example 1:
Input:
bits = [1, 0, 0]
Output: True
Explanation:
The only way to decode it is two-bit character and one-bit character.
So the last character is one-bit character.


Example 2:
Input:
bits = [1, 1, 1, 0]
Output: False
Explanation:
The only way to decode it is two-bit character and two-bit character.
So the last character is NOT one-bit character.


Note:
1 <= len(bits) <= 1000.
bits[i] is always 0 or 1.
*/

// https://leetcode.com/problems/1-bit-and-2-bit-characters/discuss/156238/Go-solutions-with-explanations
func isOneBitCharacter(bits []int) bool {
	/* if last bit is 1, return false
	   if last bit is 0, possibly true
	   from the last 0 to the second last 0: odd numbers of 1, return false
	      else return true
	*/
	if len(bits) == 0 || bits[len(bits)-1] == 1 {
		return false
	}
	if len(bits) == 1 && bits[0] == 0 {
		return true
	}

	count := 0
	l := len(bits)

	// count from tail
	for i := l - 2; i >= 0; i-- {
		if bits[i] == 1 {
			count = count + 1
		} else {
			break
		}
	}

	// if count is even return `true`, else `false`
	if count%2 == 0 {
		return true
	}

	return false
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/1-bit-and-2-bit-characters/discuss/578745/go-clean-code-0ms-beats-100

func isOneBitCharacter(bits []int) bool {
	var one int
	for i := len(bits) - 2; i >= 0 && bits[i] != 0; i-- {
		one++
	}
	return one%2 == 0
}
*/

/*
solution
- https://leetcode.com/problems/1-bit-and-2-bit-characters/discuss/558739/Go-Recursive

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	if len(bits) == 1 { // base case
		return bits[0] == 0
	}
	if bits[0] == 1 && bits[1] == 0 || // skip 2-bits number
		bits[0] == 1 && bits[1] == 1 {
		return isOneBitCharacter(bits[2:])
	}
	if bits[0] == 0 {
		return isOneBitCharacter(bits[1:])
	}
	return false
}
*/
