package main

/*
X is a good number if after rotating each digit individually by 180 degrees,
we get a valid number that is different from X.
Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves;
2 and 5 rotate to each other (on this case they are rotated in a different direction, in other words 2 or 5 gets mirrored);
6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation:

10 => 01
101 => 010
110 => 001
1001 => 0110

There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.

Note:
N  will be in range [1, 10000].
*/

// https://leetcode.com/problems/rotated-digits/discuss/579348/go-clean-code-0ms-beats-100
func rotatedDigits(N int) int {
	var cnt int
	for i := 2; i <= N; i++ {
		if isValid(i) {
			cnt++
		}
	}
	return cnt
}

func isValid(num int) bool {
	var ret bool
	for ; num > 0; num /= 10 {
		switch num % 10 {
		case 2, 5, 6, 9:
			ret = true
		case 3, 4, 7:
			return false
		}
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/rotated-digits/discuss/719263/Go-with-explanation-of-the-problem

```
The description of this problem makes us think a bit before it can be solved.
Imagine each of this digit is on a digital LED display with 7 bars (like digital clock)

Suppose we have 1234, when we rotate each digit by 180 degrees, we get 15Eh, it is invalid.
Few more numbers,

1269->1596 is valid because its different from original number
4444->hhhh === invalid
8181->8181 ==== invalid because its same as the original number after rotation
```

func Check(i int) int {
    M := make(map[int] int)
	 // digits that their own mirror
    M[0] = 0
    M[1] = 1
    M[8] = 8
	// digits that mirror each other
    M[2] = 5
    M[5] = 2
    M[6] = 9
    M[9] = 6
	// invalid digits
    M[3] = -1
    M[4] = -1
    M[7] = -1
    has256or9 := 0 // find at least one digit that is either 2,4,6,9 as this will make the number different from original
    for i > 0 {
        d := i % 10
        if M[d] < 0 { return 0}
        i = i / 10
        if M[d] != d{ has256or9 = 1}
    }
    return has256or9
}
func rotatedDigits(N int) (r int) {
    for i:=2; i <= N; i++ {
        r += Check(i)
    }
    return r
}
*/

/*
solution2:
- https://leetcode.com/problems/rotated-digits/discuss/369453/Go-golang-0ms-solution

func rotatedDigits(N int) int {

	nums := []rune{'0', '1', '5', '-', '-', '2', '9', '-', '8', '6'}
	count := 0

	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)
		same := 0
		for i, v := range s {
			if nums[v-'0'] == '-' {
				break
			}
			if nums[v-'0'] == v {
				same++
			}
			if same != len(s) && i == len(s)-1 {
				count++
			}
		}
	}

	return count

}
*/
