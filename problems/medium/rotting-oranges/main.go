package medium_rotting_oranges

type coords struct {
	i int
	j int
}

/**
https://leetcode.com/problems/rotting-oranges/
*/
func orangesRotting(grid [][]int) int {
	queue := make([]coords, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, coords{i, j})
			}
		}
	}

	minutes := 0
	for {
		nextQueue := make([]coords, 0)
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

				if x1 >= 0 && x1 < len(grid) &&
					y1 >= 0 && y1 < len(grid[0]) {
					if grid[x1][y1] == 1 {
						grid[x1][y1] = 2
						nextQueue = append(nextQueue, coords{x1, y1})
					}
				}
			}
		}
		if len(nextQueue) == 0 {
			break
		}
		minutes++
		queue = nextQueue
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
