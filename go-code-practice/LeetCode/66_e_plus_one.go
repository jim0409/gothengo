package main

import "fmt"

/*
Given a non-empty array of digits representing a non-negative integer,

plus one to the integer.

The digits are stored such that the most significant digit is

at the head of the list, and each element in the array contain

a single digit.

You may assume the integer does not contain any leading zero,

except the number 0 itself.

example 1.

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.


example 2.

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

func plusOne(digits []int) []int {
	// add 1 to the final digits ... digits[len(digits)-1]+1
	// if the digit is over 10, add 1 to the front one
	// more over, if the digit is head. decalre new array and append it with origin digits
	for i := len(digits) - 1; i >= 0; i-- {
		// if the digits[i] is less than 9, mean while the sum is less than 10
		// after add one, the digits would return
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		// otherwise the digits[i] is 9, would note with a carry
		digits[i] = 0
	}

	digits = append([]int{1}, digits...)

	return digits
}

func main() {
	input1 := []int{1, 2, 3}
	fmt.Println(plusOne(input1))

	input2 := []int{9}
	fmt.Println(plusOne(input2))
}

/*
func plusOne(digits []int) []int {
    for i := len(digits)-1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i] += 1
            return digits
        }
        digits[i] = 0
    }

    digits = append([]int{1}, digits...)

    return digits
}
*/
