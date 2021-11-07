package main

import "sort"

/*
Given an integer array nums,

return the largest perimeter of a triangle with a non-zero area,

formed from three of these lengths.

If it is impossible to form any triangle of a non-zero area, return 0.

Example 1:
Input: nums = [2,1,2]
Output: 5

Example 2:
Input: nums = [1,2,1]
Output: 0

Example 3:
Input: nums = [3,2,3,4]
Output: 10

Example 4:
Input: nums = [3,6,2,3]
Output: 8

Constraints:

3 <= nums.length <= 10^4
1 <= nums[i] <= 10^6
*/

/*
// 三角形定義，任意兩邊和大於第三邊
// 第一次寫的寫法!

func compare(a, b int, c int) bool {
	if a+b > c {
		return true
	}
	return false
}
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	maxPerimeter := 0
	for i := 0; i <= len(nums)-3; i++ {
		tmpMaxPerimeter := 0
		a := nums[i]
		b := nums[i+1]
		c := nums[i+2]
		if compare(a, b, c) && compare(b, c, a) && compare(c, a, b) {
			tmpMaxPerimeter = a + b + c
		}
		if tmpMaxPerimeter > maxPerimeter {
			maxPerimeter = tmpMaxPerimeter
		}
	}
	return maxPerimeter
}
*/

// https://leetcode.com/problems/largest-perimeter-triangle/discuss/259187/simple-golang
func largestPerimeter(A []int) int {
	sort.Ints(A)

	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

func main() {

}

/*
// quick sort solution:
refer:
- https://leetcode.com/problems/largest-perimeter-triangle/discuss/916323/golang-quicksort-solution-faster-than-100-and-memory-usage-less-than-100

func largestPerimeter(A []int) int {
	quickSort(A, 0, len(A)-1)
	for i := len(A)-1; i >= 2; i-- {
	    if (A[i] + A[i-1]) > A[i-2] && (A[i-1]+A[i-2]) > A[i] && (A[i-2] + A[i]) > A[i-1] {
		return A[i] + A[i-1] + A[i-2]
	    }
	}
	return 0
}

func quickSort(nums []int, left, right int) {
	if left >= right {
	    return
	}
	temp := nums[left]
	i, j := left, right
	for i != j {
	    for i < j && nums[j] >= temp {
		j--
	    }
	    for i < j && nums[i] <= temp {
		i++
	    }
	    if i != j {
		t := nums[i]
		nums[i] = nums[j]
		nums[j] = t
	    }
	}
	nums[left] = nums[j]
	nums[j] = temp
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
*/

/*
refer:
- https://leetcode.com/problems/largest-perimeter-triangle/discuss/1543349/golang-bubble-sort-with-shortcut-oror-beats-100
// bubble sort with shortcut

func largestPerimeter(nums []int) int {

    var largest int
    for i := 0; i < len(nums) - 1; i++ {
        largest = i
        for j := i + 1; j < len(nums); j++ {
            if nums[largest] < nums[j] {
                largest = j
            }
        }
        nums[largest], nums[i] = nums[i], nums[largest]
        if i > 1 {
            if nums[i-2] < nums[i-1] + nums[i] {
                return nums[i-2] + nums[i-1] + nums[i]
            }
        }
    }

    if nums[len(nums) - 3] < nums[len(nums) - 2] + nums[len(nums) - 1] {
        return nums[len(nums) - 3] + nums[len(nums) - 2] + nums[len(nums) - 1]
    }

    return 0
}
*/
