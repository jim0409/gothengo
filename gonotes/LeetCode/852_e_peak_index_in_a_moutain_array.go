package main

import "fmt"

/*
Let's call an array arr a mountain if the following properties hold:

- arr.length >= 3
- There exists some i with 0 < i < arr.length - 1 such that:
	arr[0] < arr[1] < ... arr[i-1] < arr[i]
	arr[i] > arr[i+1] > ... > arr[arr.length - 1]

Given an integer array arr that is guaranteed to be a mountain,
return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

Example 1:
Input: arr = [0,1,0]
Output: 1

Example 2:
Input: arr = [0,2,1,0]
Output: 1

Example 3:
Input: arr = [0,10,5,2]
Output: 1

Example 4:
Input: arr = [3,4,5,1]
Output: 2

Example 5:
Input: arr = [24,69,100,99,79,78,67,36,26,19]
Output: 2

Constraints:
- 3 <= arr.length <= 104
- 0 <= arr[i] <= 106
- arr is guaranteed to be a mountain array.
*/

// func isDecr(i1, i2 int) bool {
// 	return i1 >= i2
// }
func isIncr(i1, i2 int) bool {
	return i1 <= i2
}

func peakIndexInMountainArray(arr []int) int {
	peakLoc := 0
	l := len(arr)

	for i := 1; i < l-1; i++ {
		a := arr[i-1]
		b := arr[i]
		c := arr[i+1]
		// if isIncr(a, b) && isDecr(b, c) {
		if isIncr(a, b) && !isIncr(b, c) {
			peakLoc = i
			break
		}
	}

	return peakLoc
}

func main() {
	arr1 := []int{0, 1, 0}
	p1 := peakIndexInMountainArray(arr1)
	fmt.Println(p1)
}

/*
refer:
- https://leetcode.com/problems/peak-index-in-a-mountain-array/discuss/540595/Go-8ms-solution

func peakIndexInMountainArray(A []int) int {

    if len(A) < 3 {
        return -1
    }
    for i := 1; i < len(A)-1; i++ {

        if A[i-1] < A[i] && A[i] > A[i+1] {
            return i
        }
    }
    return -1
}
*/
