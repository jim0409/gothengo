package main

import "fmt"

/*
Given an array nums and a value val,

remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array,

you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed.

It doesn't matter what you leave beyond the new length.
*/

// 直覺寫法
func removeElement(nums []int, val int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			res--
		}
	}
	return res
}

func main() {
	testInt := []int{3, 2, 2, 3}
	testVal := 3
	fmt.Println(removeElement(testInt, testVal))
}

/*
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference,

which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeElement(nums, val);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

*/

/*
https://leetcode.com/problems/remove-element/solution/
Hints
Try two pointers.
Did you use the fact that the order of elements can be changed?
What happens when the elements to remove are rare?

參考解

func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    i, j := 0, 1
    for i < len(nums) {
        if nums[i] == val {
            for j < len(nums) && nums[j] == val {
                j++
            }
            if j >= len(nums) {
                return i
            }
            nums[i], nums[j] = nums[j], nums[i]
        }
        i++
        j++
    }
    return i

    // { 1, 1, 2, 2, 4 } , val = 1
         i  j
    -> { 2, 1, 1, 2, 4 }
            i  j
    // { 2, 2, 1, 1, 4 }
               i  j
    -> { 2, 2, 1, 1, 4 }
               i     j
    // { 2, 2, 4, 1, 1 }  val = 1
                  ^     *
                  i     j .. (j is out of length)

    做兩個指標的互換
*/
