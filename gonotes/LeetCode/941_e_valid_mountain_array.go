package main

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

e.g.
[Mountain Array]
0 2 3 4 5 2 1 0
[NOT Mountain Array]
0 2 3 3 5 2 1 0

Example 1:
Input: arr = [2,1]
Output: false

Example 2:
Input: arr = [3,5,5]
Output: false

Example 3:
Input: arr = [0,3,2,1]
Output: true


Constraints:
1 <= arr.length <= 104
0 <= arr[i] <= 104
*/

/*
1.只能有一個尖點
2.前面遞增，後面遞減
*/

func validMountainArray(A []int) bool {
	var left, right int
	for left = 0; left < len(A)-2 && A[left] < A[left+1]; left++ {
	}
	for right = len(A) - 1; right > 1 && A[right] < A[right-1]; right-- {
	}
	return left == right && left != 0
}

func main() {}

/*
refer:
- https://leetcode.com/problems/valid-mountain-array/discuss/444563/4-%2B-1-(20-ms-6.3MB)-different-solutions-in-Go


-----
1. Using left & right pointers
Runtime: 24 ms, faster than 89.92%
Memory Usage: 6.3 MB, less than 100.00%

func validMountainArray(A []int) bool {
  var left, right int
  for left = 0; left < len(A)-2 && A[left] < A[left+1]; left++ {
  }
  for right = len(A) - 1; right > 1 && A[right] < A[right-1]; right-- {
  }
  return left == right && left != 0
}


-----
2. Checking the trend
Runtime: 24 ms, faster than 89.92%
Memory Usage: 6.3 MB, less than 100.00%

func validMountainArray(A []int) bool {
  passedTop := false
  for i := len(A) - 1; i >= 2; i-- {
    switch {
    case A[i] < A[i-1] && A[i-1] < A[i-2]:
      if passedTop {
        return false
      }
    case A[i] > A[i-1] && A[i-1] > A[i-2]:
      if !passedTop {
        return false
      }
    case A[i] < A[i-1] && A[i-1] > A[i-2] && !passedTop:
      passedTop = true
    default:
      return false
    }
  }
  return passedTop
}


-----
3. Going one direction without a flag
Runtime: 24 ms, faster than 89.92%
Memory Usage: 6.3 MB, less than 100.00%

func validMountainArray(A []int) bool {
  var cursor int
  for cursor = 0; cursor < len(A)-1 && A[cursor] < A[cursor+1]; cursor++ {
  }
  if cursor == 0 || cursor == len(A)-1 {
    return false
  }
  for ; cursor < len(A)-1 && A[cursor] > A[cursor+1]; cursor++ {
  }
  return cursor == len(A)-1
}


-----
4. Going one direction with a flag
Runtime: 24 ms, faster than 89.92%
Memory Usage: 6.3 MB, less than 100.00%

func validMountainArray(A []int) bool {
  if len(A) < 3 || A[0] >= A[1] {
    return false
  }
  goingUp := true
  for i := 2; i < len(A); i++ {
    if A[i] == A[i-1] {
      return false
    }
    if goingUp {
      if A[i] < A[i-1] {
        goingUp = false
      }
    } else {
      if A[i] > A[i-1] {
        return false
      }
    }
  }
  return !goingUp
}


ex. Passed but incorrect solution

the code below was accepted at 2019-12-05T11:00:00Z
Runtime: 20 ms, faster than 97.48%
Memory Usage: 6.3 MB, less than 100.00%

func validMountainArray(A []int) bool {
  if len(A) < 3 || A[0] >= A[1] {
    return false
  }
  prev := A[1]
  goingUp := true
  for _, cur := range A[2:] {
    if goingUp {
      if cur < prev {
        goingUp = false
      }
    } else {
      if cur > prev || cur == prev {
        return false
      }
    }
    prev = cur
  }
  return !goingUp
}
*/
