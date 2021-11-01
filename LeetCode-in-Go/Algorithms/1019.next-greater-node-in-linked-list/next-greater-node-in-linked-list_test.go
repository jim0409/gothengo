package problem1019

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{225, 172, 286, 36, 70, 170, 19, 187, 206, 132, 239, 79, 58, 236, 117, 139, 278, 222, 257, 168, 152, 121, 205, 103, 130, 209, 220, 90, 263, 171, 116, 297, 165, 20, 246, 225, 205, 9, 103, 220, 14, 166, 269, 135, 214, 11, 297, 268, 214, 227, 256, 157, 106, 153, 261, 156, 165, 242, 297, 235, 297, 219, 215, 127, 258, 230, 47, 231, 7, 183, 150, 140, 288, 47, 269, 252, 41, 115, 13, 119, 273, 199, 149, 142, 175, 172, 107, 255, 53, 23, 270, 226, 146, 159, 170, 105, 5, 106, 238, 50, 210, 175, 258, 269, 272, 283, 1, 154, 269, 9, 81, 176, 252, 82, 272, 143, 66, 161, 238, 88, 251, 179, 8, 125, 37, 271, 22, 269, 268, 29, 50, 159, 151, 203, 127, 297, 275, 29, 299, 135, 219, 118, 94, 106, 89, 279, 178, 49, 159, 270, 263, 7, 85, 122, 249, 181, 276, 68, 56, 209, 128, 151, 218, 30, 70, 53, 11, 145, 153, 88, 51, 128, 101, 152, 102, 40, 115, 161, 177, 58, 276, 179, 242, 208, 269, 241, 4, 38, 217, 237, 4, 200, 189, 40, 8, 183, 196, 49, 57, 63, 157, 59, 203, 35, 269, 28, 231, 88, 272, 234, 128, 198, 122, 201, 289, 236, 1, 148, 21, 253, 111, 178, 279, 69, 122, 49, 2, 137, 221, 187, 290, 288, 296, 158, 250, 295, 44, 245, 66, 72, 129, 23, 34, 125, 288, 59, 26, 25, 135, 181, 171, 49, 80, 161, 193, 249, 284, 89, 125, 130, 81, 43, 15, 156, 228, 124, 62, 24, 83, 11, 117, 78, 146, 54, 95, 153, 189, 264, 282, 139, 82, 122, 49, 137, 88, 278, 56, 117, 105, 300, 14, 87, 243, 27, 113, 65, 30, 161, 91, 175},
		[]int{286, 286, 297, 70, 170, 187, 187, 206, 239, 239, 278, 236, 236, 278, 139, 278, 297, 257, 263, 205, 205, 205, 209, 130, 209, 220, 263, 263, 297, 297, 297, 299, 246, 246, 269, 269, 220, 103, 220, 269, 166, 269, 297, 214, 297, 297, 299, 297, 227, 256, 261, 261, 153, 261, 297, 165, 242, 297, 299, 297, 299, 258, 258, 258, 288, 231, 231, 288, 183, 288, 288, 288, 297, 269, 273, 273, 115, 119, 119, 273, 283, 255, 175, 175, 255, 255, 255, 270, 270, 270, 272, 238, 159, 170, 238, 106, 106, 238, 258, 210, 258, 258, 269, 272, 283, 297, 154, 269, 272, 81, 176, 252, 272, 272, 297, 161, 161, 238, 251, 251, 271, 271, 125, 271, 271, 297, 269, 297, 297, 50, 159, 203, 203, 297, 297, 299, 299, 299, 300, 219, 279, 279, 106, 279, 279, 289, 270, 159, 270, 276, 276, 85, 122, 249, 276, 276, 289, 209, 209, 218, 151, 218, 276, 70, 145, 145, 145, 153, 161, 128, 128, 152, 152, 161, 115, 115, 161, 177, 276, 276, 289, 242, 269, 269, 272, 269, 38, 217, 237, 269, 200, 203, 196, 183, 183, 196, 203, 57, 63, 157, 203, 203, 269, 269, 272, 231, 272, 272, 289, 289, 198, 201, 201, 289, 290, 253, 148, 253, 253, 279, 178, 279, 290, 122, 137, 137, 137, 221, 290, 290, 296, 296, 300, 250, 295, 300, 245, 288, 72, 129, 288, 34, 125, 288, 300, 135, 135, 135, 181, 193, 193, 80, 161, 193, 249, 284, 300, 125, 130, 156, 156, 156, 156, 228, 264, 146, 83, 83, 117, 117, 146, 146, 153, 95, 153, 189, 264, 282, 300, 278, 122, 137, 137, 278, 278, 300, 117, 300, 300, 0, 87, 243, 0, 113, 161, 161, 161, 175, 175, 0},
	},

	{
		[]int{2, 1, 5},
		[]int{5, 5, 0},
	},

	{
		[]int{2, 7, 4, 3, 5},
		[]int{7, 0, 5, 5, 0},
	},

	{
		[]int{1, 7, 5, 1, 9, 2, 5, 1},
		[]int{7, 9, 9, 9, 0, 5, 0, 0},
	},

	// 可以有多个 testcase
}

func Test_nextLargerNodes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		head := kit.Ints2List(tc.head)
		ast.Equal(tc.ans, nextLargerNodes(head), "输入:%v", tc)
	}
}

func Benchmark_nextLargerNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := kit.Ints2List(tc.head)
			nextLargerNodes(head)
		}
	}
}
