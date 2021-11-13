package main

import "fmt"

/*
Given two sorted integer arrays nums1 and nums2,

merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.

You may assume that nums1 has enough space
(size that is greater or equal to m + n) to hold additional elements from nums2.

Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

// nums1以及nums2均為有序數列
func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1有值的最後一位數
	i := m - 1
	j := n - 1
	// k 為num1這個陣列的最後一位數
	k := m + n - 1
	// 重新規劃nums1
	// i 是nums1的最後一位數
	// j 是nums2的最後一位數
	for i >= 0 && j >= 0 {
		// 如果nums1的最後一位數大於nums2的最後一位數就
		// 將nums1的k位數替換為nums1的最後一位數，並且將最後一位數遞減1
		// 反之操作nums2
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		// 再賦予最後一位數的值後，遞減要賦予的k位數
		k--
	}

	// edge case，如果比較完畢後nums2的值相較於nums1的值小，則將nums2剩下的值塞滿nums1值的前面
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

}

/*
Hint1:
You can easily solve this problem if you simply think about two elements at a time
rather than two arrays.

We know that each of the individual arrays is sorted.

What we don't know is how they will intertwine.

Can we take a local decision and arrive at an optimal solution?


Hint2:
If you simply consider one element each at a time from the two arrays

and make a decision and proceed accordingly, you will arrive at the optimal solution.
*/

/*
solution:
- https://leetcode.com/problems/merge-sorted-array/discuss/509594/Go


func merge(nums1 []int, m int, nums2 []int, n int)  {

    i, m, n := len(nums1) - 1, m - 1, n - 1

    for m >= 0 && n >= 0 {
        if nums1[m] > nums2[n] {
            nums1[i] = nums1[m]
            m--
        } else {
            nums1[i] = nums2[n]
            n--
        }
        i--
    }
    for m >= 0 {
        nums1[i] = nums1[m]
        m--
        i--
    }
    for n >= 0 {
        nums1[i] = nums2[n]
        n--
        i--
    }
    nums1 = nums1[i + 1:]
    return
}
*/
