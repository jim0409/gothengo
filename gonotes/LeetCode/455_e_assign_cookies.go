package main

import "sort"

/*
Assume you are an awesome parent and want to give your children some cookies.
But, you should give each child at most one cookie.
Each child `i` has a greed factor `gi`, which is the minimum size of a cookie that the child will be content with;
and each cookie `j` has a size `sj`. If sj >= gi, we can assign the cookie j to the child i,
and the child i will be content.

Your goal is to maximize the number of your content children and output the maximum number.

Note:
You may assume the greed factor is always positive.
You cannot assign more than one cookie to one child.

Example 1:
Input: [1,2,3], [1,1]

Output: 1

Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3.
And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
You need to output 1.


Example 2:
Input: [1,2], [1,2,3]

Output: 2

Explanation: You have 2 children and 3 cookies. The greed factors of 2 children are 1, 2.
You have 3 cookies and their sizes are big enough to gratify all of the children,
You need to output 2.
*/

func findContentChildren(g []int, s []int) (i int) {
	/*
		g(greed) 為小孩的餅乾期望值, s(satisfy)表示餅乾能滿足的程度
		先將兩者做排序，盡量從最小的開始滿足
	*/
	sort.Ints(g)
	sort.Ints(s)
	for j := 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			i++
		}
	}
	return
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/assign-cookies/discuss/466312/go-clean-code-24ms-beats-100


func findContentChildren(g []int, s []int) (i int) {
	sort.Ints(g)
	sort.Ints(s)
	for j := 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			i++
		}
	}
	return
}
*/
