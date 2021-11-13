package main

import "fmt"

/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array),
some elements appear twice and others appear once.
Find all the elements of [1, n] inclusive that do not appear in this array.

* Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]

可以將其排序為
1,2,3,4,7,8 => 缺少 5 & 6
*/

// // declare extra space ... O(n) memory ... but faster...
// func findDisappearedNumbers(a []int) []int {
// 	for i := range a {
// 		for a[a[i]-1] != a[i] {
// 			a[a[i]-1], a[i] = a[i], a[a[i]-1]
// 		}
// 	}

// 	out := []int{}
// 	for i, v := range a {
// 		if v != i+1 {
// 			out = append(out, i+1)
// 		}
// 	}
// 	return out
// }

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		// 如果剛好 i 的順序 等於 nums[i] ... continue
		if i+1 == nums[i] || nums[i] == 0 {
			continue
		}

		// correct value already in swap place ... 如果剛好 nums[i] 已經存在過了。將nums[i]宣告為0，後面可以刪除
		if nums[i] == nums[nums[i]-1] {
			nums[i] = 0
			continue
		}

		// swap ... 假設 nums[i]不為i+1，不為0，且也沒有跟其他對應的數字重複過。將nums[i]與nums[nums[i]-1]做互換
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

		// redo this spot
		i--
	}

	// fmt.Println(nums) // [1 2 3 4 0 0 7 8]

	// convert existing numbers to zeros and add missing numbers ... 將存在的數字轉為 0 後加入其他不存在的數字
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = i + 1
			continue
		}

		nums[i] = 0
	}

	// delete zeros
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return nums
}

func main() {
	input1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(input1))
}

/*
solution:
- https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/discuss/317430/Go-100-memory-O(1)-memory


func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
        if i + 1 == nums[i] || nums[i] == 0 {
            continue
        }

        // correct value already in swap place
        if nums[i] == nums[nums[i] - 1] {
            nums[i] = 0
            continue
        }

        // swap
        nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]

        // redo this spot
        i--
    }

    // convert existing numbers to zeros and add missing numbers
    for i := 0; i < len(nums); i++ {
        if (nums[i] == 0) {
            nums[i] = i + 1
            continue
        }

        nums[i] = 0
    }

    // delete zeros
    for i := 0; i < len(nums); i++ {
        if (nums[i] == 0) {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }

    return nums
}
*/
