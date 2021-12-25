package easy_flood_fill

/**
https://leetcode.com/problems/flood-fill/
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	fromColor := image[sr][sc]
	if newColor == fromColor {
		return image
	}
	fill(image, sr, sc, fromColor, newColor)

	return image
}

func fill(image [][]int, x int, y int, fromColor int, toColor int) {
	if image[x][y] == fromColor {
		image[x][y] = toColor
		if x > 0 {
			fill(image, x-1, y, fromColor, toColor)
		}
		if y > 0 {
			fill(image, x, y-1, fromColor, toColor)
		}
		if x+1 < len(image) {
			fill(image, x+1, y, fromColor, toColor)
		}
		if y+1 < len(image[0]) {
			fill(image, x, y+1, fromColor, toColor)
		}
	}
}
