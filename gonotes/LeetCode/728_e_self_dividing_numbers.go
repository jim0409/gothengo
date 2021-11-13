package main

/*
A self-dividing number is a number that is divisible by every digit it contains.

For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

Also, a self-dividing number is not allowed to contain the digit zero.

Given a lower and upper number bound,
output a list of every possible self dividing number, including the bounds if possible.

Example 1:
Input:
left = 1, right = 22
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]

Note:
The boundaries of each input argument are 1 <= left <= right <= 10000.
*/

func selfDividingNumbers(left int, right int) []int {
	// 宣告要回傳的 int 陣列
	res := []int{}

	// 當左數小於右數時，持續迭代
	for left <= right {
		// 讓左數為`v`
		v := left

		// 宣告一個暫存數，並且持續迭代。判斷`v`是否符合自除數的規則(128 => /1 /2 /8 餘數皆為 0)
		tmp := v
		for v > 0 {
			d := v % 10
			if d == 0 || (tmp%d) != 0 {
				break
			}
			v /= 10
		}

		if v == 0 {
			// 如果餘數為 0，將 tmp 放入回傳結果
			res = append(res, tmp)
		}
		// 左數遞進1
		left++
	}

	return res
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/self-dividing-numbers/discuss/578798/go-clean-code-0ms-beats-100

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0, right-left+1)
	for i := left; i <= right; i++ {
		if isValid(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func isValid(num int) bool {
	for n := num; n > 0; n /= 10 {
		if n%10 == 0 || num%(n%10) != 0 {
			return false
		}
	}
	return true
}
*/

/*
solution2:
- https://leetcode.com/problems/self-dividing-numbers/discuss/465245/Go-double-100

func selfDividingNumbers(left int, right int) []int {
	output := []int{}
loop:
	for i := left; i <= right; i++ {
		number := i
		for number != 0 {
            if (number%10) == 0 || i%(number%10) != 0 {
				continue loop
			}
			number /= 10
		}
		output = append(output, i)
	}
	return output
}
*/
