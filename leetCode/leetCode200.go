package leetCode

import "fmt"

// 200. 岛屿数量

func TestLeetCode200() {
	grids1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grids1))
}

func numIslands(grid [][]byte) int {
	// 相连的1算1个
	// 深度遍历 访问过的标记
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs200(grid, i, j)
			}
		}
	}
	return count
}

// 深度搜索 遇到水或访问过的停止
func dfs200(grid [][]byte, i, j int) {
	// 上下左右的访问
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '*'
	dfs200(grid, i-1, j)
	dfs200(grid, i+1, j)
	dfs200(grid, i, j-1)
	dfs200(grid, i, j+1)
}
