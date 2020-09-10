package Problem

//leetcode submit region begin(Prohibit modification and deletion)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		dfsFloodFill(image, sr, sc, newColor, color)
	}
	return image
}

func dfsFloodFill(image [][]int, x, y, newColor, color int) {
	if color == image[x][y] {
		image[x][y] = newColor
		if x+1 < len(image) {
			dfsFloodFill(image, x+1, y, newColor, color)
		}
		if x-1 >= 0 {
			dfsFloodFill(image, x-1, y, newColor, color)
		}
		if y+1 < len(image[0]) {
			dfsFloodFill(image, x, y+1, newColor, color)
		}
		if y-1 >= 0 {
			dfsFloodFill(image, x, y-1, newColor, color)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

// 2020-08-16 09:02:55

//
//An image is represented by a 2-D array of integers, each integer representing
//the pixel value of the image (from 0 to 65535).
//
//Given a coordinate (sr, sc) representing the starting pixel (row and column) o
//f the flood fill, and a pixel value newColor, "flood fill" the image.
//
//To perform a "flood fill", consider the starting pixel, plus any pixels connec
//ted 4-directionally to the starting pixel of the same color as the starting pixe
//l, plus any pixels connected 4-directionally to those pixels (also with the same
// color as the starting pixel), and so on. Replace the color of all of the aforem
//entioned pixels with the newColor.
//
//At the end, return the modified image.
//
// Example 1:
//
//Input:
//image = [[1,1,1],[1,1,0],[1,0,1]]
//sr = 1, sc = 1, newColor = 2
//Output: [[2,2,2],[2,2,0],[2,0,1]]
//Explanation:
//From the center of the image (with position (sr, sc) = (1, 1)), all pixels con
//nected
//by a path of the same color as the starting pixel are colored with the new col
//or.
//Note the bottom corner is not colored 2, because it is not 4-directionally con
//nected
//to the starting pixel.
//
//
//
// Note:
// The length of image and image[0] will be in the range [1, 50].
// The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < im
//age[0].length.
// The value of each color in image[i][j] and newColor will be an integer in [0,
// 65535].
//
// 👍 105 👎 0
