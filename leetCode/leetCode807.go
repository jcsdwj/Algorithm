/**
  @author: 伍萬
  @since: 2024/7/14
  @desc: //TODO
**/

package leetCode

import "fmt"

//807. 保持城市天际线

func TestLeetCode807() {
	grid1 := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid1))
	grid2 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid2))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// 求每行每列的最大值
	dpRow := make([]int, m)    // 保存行最大值
	dpColumn := make([]int, n) // 保存列最大值

	// 求每行最大值
	for i := 0; i < m; i++ {
		row := grid[i][0]
		for j := 0; j < n; j++ {
			if grid[i][j] > row {
				row = grid[i][j]
			}
		}
		dpRow[i] = row
	}

	// 求每列最大值
	for j := 0; j < n; j++ {
		column := grid[0][j]
		for i := 0; i < m; i++ {
			if grid[i][j] > column {
				column = grid[i][j]
			}
		}
		dpColumn[j] = column
	}

	sum := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := min(dpColumn[j], dpRow[i])
			sum += max(temp-grid[i][j], 0) // 保证大于0
		}
	}

	return sum
}
