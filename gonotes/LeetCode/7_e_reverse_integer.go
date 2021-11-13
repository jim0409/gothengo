package main

import (
	"fmt"
)

/*
	given a 32-bit signed integer, reverse digits of an integer
	Input 123
	output: 321
*/
/*
	Note:
	Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
	For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func main() {
	a := 321
	fmt.Println(reverse(a))

	b := 10
	fmt.Println(reverse(b))

	c := -2147483648
	fmt.Println(reverse(c))
}

func reverse(x int) int {
	var answer int
	tmpInt := x

	for tmpInt != 0 {
		answer = answer*10 + tmpInt%10
		tmpInt = tmpInt / 10
	}

	if answer > 0x7fffffff || answer < -1*0x80000000 {
		return 0
	}

	return answer
}

/*
func reverse(x int) int {
	var tmpInt int

	sign := 1
	if x < 0 {
		sign = -1
	}
	tmpInt = sign * x
	reverseIntArr := []int{}
	var result int

	for tmpInt >= 10 {
		curInt := tmpInt % 10
		reverseIntArr = append(reverseIntArr, curInt)
		tmpInt = int(tmpInt / 10)
	}
	if tmpInt > 0 {
		reverseIntArr = append(reverseIntArr, tmpInt)
	}

	fmt.Println(len(reverseIntArr))

	count := 1
	for i := len(reverseIntArr) - 1; i >= 0; i-- {
		result += reverseIntArr[i] * count
		count = count * 10
	}

	if result*sign > 2147483648-1 || result*sign < -2147483648 {
		return 0
	} else {
		return result * sign
	}

}
*/
