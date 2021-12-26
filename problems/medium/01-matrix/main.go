package medium_01_matrix

type coords struct {
	i int
	j int
}

/**
https://leetcode.com/problems/01-matrix/
*/
func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}

	distMat := make([][]int, len(mat))
	queue := make([]coords, 0)
	for i := 0; i < len(mat); i++ {
		distMat[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				distMat[i][j] = -1
			} else {
				queue = append(queue, coords{i, j})
			}
		}
	}

	for len(queue) > 0 {
		x, y := queue[0].i, queue[0].j
		queue = queue[1:]

		for r := 0; r < 4; r++ {
			x1 := x - 1
			y1 := y
			switch r {
			case 1:
				x1, y1 = x, y-1
			case 2:
				x1, y1 = x+1, y
			case 3:
				x1, y1 = x, y+1
			}

			if x1 >= 0 && x1 < len(mat) &&
				y1 >= 0 && y1 < len(mat[0]) {
				if distMat[x1][y1] == -1 {
					distMat[x1][y1] = distMat[x][y] + 1
					queue = append(queue, coords{x1, y1})
				}
			}
		}
	}

	return distMat
}
