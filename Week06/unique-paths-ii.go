package main

import "fmt"

//63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	f := make([]int, len(obstacleGrid))
	if obstacleGrid[0][0] == 0 {
		return 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}

		}
	}
	return f[len(f)-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}))
}
