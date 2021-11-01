package main

/*
Given an array A of non-negative integers,
return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.

Example 1:
Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.


Note:
1 <= A.length <= 5000
0 <= A[i] <= 5000
*/

func sortArrayByParity(A []int) []int {
	for low, high := 0, len(A)-1; low < high; {
		for ; (A[low]&1 == 0) && low < high; low++ {
		}
		for ; (A[high]&1 == 1) && low < high; high-- {
		}
		A[low], A[high] = A[high], A[low]
	}
	return A
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/sort-array-by-parity/discuss/343124/Go-golang-clean-solution

func sortArrayByParity(A []int) []int {
    for left, right := 0, 0; right < len(A); right++ {
        if A[right] % 2 == 0 { A[left], A[right] = A[right], A[left]; left++ }
    }
    return A
}

similar to
- https://leetcode.com/problems/sort-array-by-parity/discuss/801186/Go-in-place

func sortArrayByParity(nums []int) []int {
    // target idx for even numbers
    nextEvenPosition := 0
    for cur := 0; cur < len(nums); cur++ {
        if nums[cur]%2 == 0 {
            nums[nextEvenPosition], nums[cur] = nums[cur], nums[nextEvenPosition]
            nextEvenPosition++
        }
    }
    return nums
}
*/
