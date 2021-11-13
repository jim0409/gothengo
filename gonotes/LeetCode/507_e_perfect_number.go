package main

import "math"

/*
We define the Perfect Number is a positive integer that is
equal to the sum of all its positive divisors except itself.

Now, given an integer n,
write a function that returns true when it is a perfect number
and false when it is not.


Example:
Input: 28
Output: True
Explanation: 28 = 1 + 2 + 4 + 7 + 14

Note: The input number n will not exceed 100,000,000. (1e8)
*/

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum, sqrt := 1, int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i != 0 {
			continue
		}
		sum += i
		if i != num/i {
			sum += num / i
		}
	}
	return sum == num
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/perfect-number/discuss/415760/Go-golang-one-line-solution

func checkPerfectNumber(num int) bool {
    return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}
*/

/*
[without sqrt] solution:
- https://leetcode.com/problems/perfect-number/discuss/341652/Simple-Go-Solution-Without-Sqrt

func checkPerfectNumber(num int) bool {
    var result int

    if num <= 1  {
        return false
    }
    i := 2
    for(i < num/i && i != num) {
        if num%i == 0 {
            result += i
            result += num/i
        }
        i++
    }

    return (result + 1) == num
}
*/
