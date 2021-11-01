package main

import "fmt"

/*
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].
An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

Example 1:
Input: [1,2,2,3]
Output: true

Example 2:
Input: [6,5,4,4]
Output: true

Example 3:
Input: [1,3,2]
Output: false

Example 4:
Input: [1,2,4,5]
Output: true

Example 5:
Input: [1,1,1]
Output: true

Note:
1 <= A.length <= 50000
-100000 <= A[i] <= 100000
*/

// func isMonotonic(A []int) bool {
// 	if A == nil {
// 		return true
// 	}

// 	// fmt.Println(A[0] >= A[1])
// 	if A[0] >= A[1] {
// 		tmp := A[1]
// 		for i := 2; i < len(A); i++ {
// 			// fmt.Println(A[i])
// 			if A[i] <= tmp {
// 				tmp = A[i]
// 			} else {
// 				return false
// 			}
// 		}
// 	} else {
// 		tmp := A[1]
// 		for i := 2; i < len(A); i++ {
// 			// fmt.Println(A[i])
// 			if A[i] >= tmp {
// 				tmp = A[i]
// 			} else {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	var diff int
	for i := 1; i < len(A); i++ {
		curr, prev := A[i], A[i-1]

		pairDiff := curr - prev
		if diff == 0 {
			diff = pairDiff
		} else if diff > 0 && pairDiff < 0 {
			return false
		} else if diff < 0 && pairDiff > 0 {
			return false
		}
	}

	return true
}

func main() {
	input1 := []int{1, 2, 2, 3}
	fmt.Println(isMonotonic(input1))

	input2 := []int{1, 1, 2}
	fmt.Println(isMonotonic(input2))

	input3 := []int{1, 1, 0}
	fmt.Println(isMonotonic(input3))
}

/*
solution:
- https://leetcode.com/problems/monotonic-array/discuss/243684/100-use-go

func increas (A []int) bool {
    for i:=0 ; i <len(A)-1; i++ {
        if A[i] > A[i+1] {
            return false
        }
    }
    return true
}
func decreas (A []int) bool {
    for i:=0 ; i <len(A)-1; i++ {
        if A[i] < A[i+1] {
            return false
        }
    }
    return true

}
func isMonotonic(A []int) bool {
    for i := 0; i<len(A)-1; i++ {
        if A[i] == A[i+1] {
            continue
        }

        if A[i] > A[i+1] {
            return decreas(A)
        }else {
            break
        }
    }
    return increas(A)
}
*/
