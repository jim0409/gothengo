package main

/*
An image is represented by a 2-D array of integers,
each integer representing the pixel value of the image (from 0 to 65535).

Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill,
and a pixel value newColor, "flood fill" the image.

To perform a "flood fill", consider the starting pixel,

1.
plus `any pixels connected 4-directionally` to the starting pixel
of the same color as the starting pixel,
2.
plus any pixels connected 4-directionally to those pixels
(also with the same color as the starting pixel), and so on.

Replace the color of all of the aforementioned pixels with the newColor.

At the end, return the modified image.


Example 1:
Input:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation:
From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
by a path of the same color as the starting pixel are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected
to the starting pixel.


Note:
The length of image and image[0] will be in the range [1, 50].
The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
The value of each color in image[i][j] and newColor will be an integer in [0, 65535].
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		helper(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func helper(image [][]int, sr, sc, oldColor, newColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor
	helper(image, sr-1, sc, oldColor, newColor)
	helper(image, sr+1, sc, oldColor, newColor)
	helper(image, sr, sc-1, oldColor, newColor)
	helper(image, sr, sc+1, oldColor, newColor)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/flood-fill/discuss/386578/Go%3A-DFS


func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {


    formerColor := image[sr][sc]

    if newColor == formerColor {
        return image
    }

    image[sr][sc] = newColor

    if hasSameColor(image, sr + 1, sc, formerColor) {
        floodFill(image, sr + 1, sc, newColor)
    }
    if hasSameColor(image, sr - 1, sc, formerColor) {
        floodFill(image, sr - 1, sc, newColor)
    }
    if hasSameColor(image, sr, sc + 1, formerColor) {
        floodFill(image, sr, sc + 1, newColor)
    }
    if hasSameColor(image, sr, sc - 1, formerColor) {
        floodFill(image, sr, sc - 1, newColor)
    }

    return image

}


func hasSameColor(image [][]int, sr, sc, formerColor int) bool {
    return sr >= 0 && sr < len(image) && sc >= 0 && sc < len(image[sr]) && image[sr][sc] == formerColor
}
*/
