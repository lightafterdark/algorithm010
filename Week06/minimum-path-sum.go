package main

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] = grid[i][j-1] + grid[i][j]
				} else {
					grid[i][j] = grid[i-1][j] + grid[i][j]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main()  {
	
}