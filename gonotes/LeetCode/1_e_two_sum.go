package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for ii := i + 1; ii < len(nums); ii++ {
			if j+nums[ii] == target {
				return []int{i, ii}
			}
		}
	}
	return nil
}

func main() {
	// given nums = [2,7,11,15]
	// targe = 9
	// return [0,1] because 2+7=9

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

	nums2 := []int{3, 2, 4, 6}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))
}

// func twoSum(nums []int, target int) []int {
//     complementMap := map[int]int{}
//     for i, n := range nums {
//         c := target - n
//         if j, ok := complementMap[c]; ok {
//             return []int{j, i}
//         }
//         complementMap[n] = i
//     }
//     return []int{}
// }
