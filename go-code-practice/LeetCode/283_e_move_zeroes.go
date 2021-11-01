package main

/*
Given an array nums,
write a function to move all 0's to the end of it
while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/move-zeroes/discuss/557813/Go


func moveZeroes(nums []int)  {
    cnt := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[cnt] = nums[i]
            cnt++
        }
    }

    for i := cnt; i < len(nums); i++ {
        nums[i] = 0
    }
}
*/
