package main

import "math"

/*
For a web developer, it is very important to know how to design a web page's size.
So, given a specific rectangular web page’s area, your job by now is to design a rectangular web page,
whose length L and width W satisfy the following requirements:

1. The area of the rectangular web page you designed must equal to the given target area.
2. The width W should not be larger than the length L, which means L >= W.
3. The difference between length L and width W should be as small as possible.

You need to output the length L and the width W of the web page you designed in sequence.

Example:
Input: 4
Output: [2, 2]
Explanation: The target area is 4, and all the possible ways to construct it are [1,4], [2,2], [4,1].
But according to requirement 2, [1,4] is illegal;
according to requirement 3,  [4,1] is not optimal compared to [2,2]. So the length L is 2, and the width W is 2.

Note:
The given area won't exceed 10,000,000 and is a positive integer
The web page's width and length you designed must be positive integers.
*/

// https://leetcode.com/problems/construct-the-rectangle/discuss/467968/go-clean-code-0ms-beats-100
func constructRectangle(area int) []int {
	for m := int(math.Sqrt(float64(area))); ; m-- {
		if area%m == 0 {
			return []int{area / m, m}
		}
	}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/construct-the-rectangle/discuss/382312/Go-golang-2-solutions


// solution1.
func constructRectangle(area int) []int {

	tmp := [][]int{}

	for i := area; i > 0; i-- {
		temp := []int{}
		if area%i == 0 && i >= area/i {
			temp = append(temp, i)
			temp = append(temp, area/i)
			tmp = append(tmp, temp)
		}
	}

	return tmp[len(tmp)-1]

}


// solution2.
// Because L >= W, find largest W
func constructRectangle(area int) []int {

	w := int(math.Sqrt(float64(area)))

	for area%w != 0 {
		w--
	}

	return []int{area/w, w}

}
*/
