package main

import "strconv"

/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz”
但，當數字是具有3的倍數的時候，打印"Fizz"

instead of the number and for the multiples of five output “Buzz”.
當數字具有5的倍數的時候，打印"Buzz"

For numbers which are multiples of both three and five output “FizzBuzz”.
如果數字同時為3也為5的倍數，打印"FizzBuzz"

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var s string
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		res = append(res, s)
	}
	return res
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/fizz-buzz/discuss/459413/go-clean-code-0ms-beats-100


func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			ret[i-1] = "FizzBuzz"
		case i%3 == 0:
			ret[i-1] = "Fizz"
		case i%5 == 0:
			ret[i-1] = "Buzz"
		default:
			ret[i-1] = strconv.Itoa(i)
		}
	}
	return ret
}
*/
