package main

import "fmt"

/*
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array,

and it should return false if every element is distinct.

Example 1:
Input: [1,2,3,1]
Output: true

Example 2:
Input: [1,2,3,4]
Output: false

Example 3:
Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

/*
// runtime 24ms - memory 7.7MB
func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _,j:=range nums{
		if m[j]{
			return true
		}
		m[j]=true
	}
	return false
}
*/

// runtime 20ms - memory 7.4MB
func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := set[v]; ok {
            return true
        } else {
            set[v] = struct{}{}
        }
    }
    return false
}

func main(){
	input1 := []int{1,2,3,1}
	fmt.Println(containsDuplicate(input1))
}


/*
solution:
- https://leetcode.com/problems/contains-duplicate/discuss/538862/Go


func containsDuplicate(nums []int) bool {
    set := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := set[v]; ok {
            return true
        } else {
            set[v] = struct{}{}
        }
    }
    return false
}
*/