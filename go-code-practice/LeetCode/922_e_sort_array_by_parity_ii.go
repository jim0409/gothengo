package main

/*
Given an array A of non-negative integers,
half of the integers in A are odd, and half of the integers are even.

Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.

You may return any answer array that satisfies this condition.


Example 1:
Input: [4,2,5,7]
Output: [4,5,2,7]
Explanation:
[4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.


Note:
2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000
*/

// https://leetcode.com/problems/sort-array-by-parity-ii/discuss/553453/go-clear-code
// "Two Pointer Solution"
func sortArrayByParityII(A []int) []int {
	i, j := 0, 1
	for i < len(A) && j < len(A) {
		if A[i]%2 == 1 {
			A[i], A[j] = A[j], A[i]
			j += 2
			continue
		}
		i += 2
	}
	return A
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/sort-array-by-parity-ii/discuss/346113/Go-golang-300ms-clean-solution

func sortArrayByParityII(A []int) []int {

	length := len(A)
	even := []int{}
	odd := []int{}
	sorted := []int{}

	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	for i := 0; i < length; i++ {
		switch i % 2 {
		case 0:
			sorted = append(sorted, even[0])
			even = even[1:]
		case 1:
			sorted = append(sorted, odd[0])
			odd = odd[1:]
		}
	}
	return sorted
}
*/
