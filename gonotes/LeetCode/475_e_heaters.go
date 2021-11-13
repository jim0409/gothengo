package main

import (
	"math"
	"sort"
)

/*
Winter is coming! Your first job during the contest is to design a standard heater
with fixed warm radius to warm all the houses.

Now, you are given positions of houses and heaters on a horizontal line,
find out minimum radius of heaters so that all houses could be covered by those heaters.

So, your input will be the positions of houses and heaters seperately,
and your expected output will be the minimum radius standard of heaters.

給定一組屋子的所在位址，以及屋子的間距
給定一組暖爐放置的位置。求加熱範圍以期所有房子都能被暖爐加熱到

Note:
Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
As long as a house is in the heaters' warm radius range, it can be warmed.
All the heaters follow your radius standard and the warm radius will the same.


Example 1:
Input: [1,2,3],[2]
[house, house, house]
[  X  , heater,  X  ]

Output: 1
Explanation: The only heater was placed in the position 2,
and if we use the radius 1 standard, then all the houses can be warmed.


Example 2:
Input: [1,2,3,4],[1,4]
[house, house, house, house]
[heater,  X  ,   X  ,heater]

Output: 1
Explanation: The two heater was placed in the position 1 and 4.
We need to use radius 1 standard, then all the houses can be warmed.
*/
func findRadius(houses []int, heaters []int) int {
	// 將房子 與 加熱器 都先做 排序 (小->大)
	sort.Ints(houses)
	sort.Ints(heaters)

	var radius int
	// i 為房子的索引，j 為加熱器的位置
	for i, j := 0, 0; i < len(houses); i++ {
		// 當 j < len(heaters)-1 : 加熱器充足，且
		// heaters[j]+heaters[j+1] <= 2*house[i] : 兩個加熱器之間的距離，小於兩間房子的距離時
		for ; j < len(heaters)-1 && heaters[j]+heaters[j+1] <= 2*houses[i]; j++ {
		}
		// 半徑為 計有半徑 與 heaters[j]-house[i]中取最大值: 加熱器的半徑最小需要涵括到鄰近的房屋
		radius = int(math.Max(float64(radius), math.Abs(float64(heaters[j]-houses[i]))))
	}
	return radius
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/heaters/discuss/467721/go-clean-code-40ms-beats-100


func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var radius int
	for i, j := 0, 0; i < len(houses); i++ {
		for ; j < len(heaters)-1 && heaters[j]+heaters[j+1] <= 2*houses[i]; j++ {}
		radius = int(math.Max(float64(radius), math.Abs(float64(heaters[j]-houses[i]))))
	}
	return radius
}
*/

/*
solution:
- https://leetcode.com/problems/heaters/discuss/95929/Binary-Search-Solution-in-Go-(beats-94)


func findRadius(houses []int, heaters []int) int {
	// 先對 headers 做排序
	sort.Ints(heaters)

	//
    res := -math.MaxInt32
	var dist1, dist2, index int


	// 將房子的位置與heaters的位置一起做排序
    for _, house := range houses {
        index = binarySearch(house, heaters)
        if index < 0 {
            index = -(index+1)
        }

        dist1 = math.MaxInt32
        if index - 1 >= 0 {
            dist1 = house - heaters[index-1]
		}

        dist2 = math.MaxInt32
        if index < len(heaters) {
            dist2 = heaters[index] - house
        }

        res = Max(res, Min(dist1, dist2))
    }
    return res
}

func Min(i int, j int) int {
    if i <= j {
        return i
    }
    return j
}

func Max(i int, j int) int {
    if i >= j {
        return i
    }
    return j
}

func binarySearch(num int, nums []int) int {
    l, r := 0, len(nums)-1
    var mid int
    for l <= r {
        mid = (l + r)/2
        if nums[mid] == num {
            return mid
        } else if nums[mid] > num{
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return l
}
*/
