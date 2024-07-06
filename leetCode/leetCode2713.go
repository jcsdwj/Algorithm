/**
  @author: 伍萬
  @since: 2024/6/19
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"sort"
)

// 2713. 矩阵中严格递增的单元格数

func TestLeetCode2713() {
	mat1 := [][]int{{3, 1}, {3, 4}}
	mat2 := [][]int{{1, 1}, {1, 1}}
	mat3 := [][]int{{3, 1, 6}, {-9, 5, 7}}
	fmt.Println(maxIncreasingCells(mat1))
	fmt.Println(maxIncreasingCells(mat2))
	fmt.Println(maxIncreasingCells(mat3))
}

// 同一行同一列
// 关键是怎么遍历
func maxIncreasingCells(mat [][]int) int {
	length := len(mat)
	if length == 0 {
		return 0
	}
	width := len(mat[0])
	if width == 0 {
		return 0
	}

	dp2713 := make([][]int, length)
	for i := 0; i < length; i++ {
		dp2713[i] = make([]int, width)
	}

	valMax := 0
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			val := dfs2713(mat, i, j, dp2713)
			if valMax < val {
				valMax = val
			}
		}
	}
	return valMax
}

// 理解错误 不是相邻节点 同行同列都行
func dfs2713(mat [][]int, curi, curj int, dp2713 [][]int) int {

	if dp2713[curi][curj] != 0 { // 如果有最长路径直接返回
		return dp2713[curi][curj]
	}

	valLeft, valRight, valUp, valDown := 0, 0, 0, 0
	val := mat[curi][curj]

	// 找比我小的最长路径
	if curi-1 >= 0 && val < mat[curi-1][curj] {
		// 判断有没有路径（也就是访问过）
		if dp2713[curi-1][curj] == 0 {
			valLeft = dfs2713(mat, curi-1, curj, dp2713)
		} else {
			valLeft = dp2713[curi-1][curj]
		}
	}
	if curi+1 < len(mat) && val < mat[curi+1][curj] {
		if dp2713[curi+1][curj] == 0 {
			valRight = dfs2713(mat, curi+1, curj, dp2713)
		} else {
			valRight = dp2713[curi+1][curj]
		}
	}
	if curj-1 >= 0 && val < mat[curi][curj-1] {
		if dp2713[curi][curj-1] == 0 {
			valUp = dfs2713(mat, curi, curj-1, dp2713)
		} else {
			valUp = dp2713[curi][curj-1]
		}
	}
	if curj+1 < len(mat[0]) && val < mat[curi][curj+1] {
		if dp2713[curi][curj+1] == 0 {
			valDown = dfs2713(mat, curi, curj+1, dp2713)
		} else {
			valDown = dp2713[curi][curj+1]
		}
	}

	dp2713[curi][curj] = 1 + max(valLeft, max(valRight, max(valUp, valDown))) // 以[curi][curj] 为终点的最长路径
	return 1 + max(valLeft, max(valRight, max(valUp, valDown)))
}

// 动态规划 dp[i][j]为移动到单元格 (i,j) 的最大步数
func maxIncreasingCells2(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	mp := make(map[int][][2]int) // 某个值对应的坐标（相同值为多个坐标）
	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[mat[i][j]] = append(mp[mat[i][j]], [2]int{i, j})
		}
	}
	keys := []int{}
	for key, _ := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		pos := mp[key]
		res := []int{} // 存放相同数值的答案，便于后续更新 row 和 col
		for _, arr := range pos {
			res = append(res, max(row[arr[0]], col[arr[1]])+1)
		}
		for i := 0; i < len(pos); i++ {
			arr, d := pos[i], res[i]
			row[arr[0]] = max(row[arr[0]], d)
			col[arr[1]] = max(col[arr[1]], d)
		}
	}

	return maxSlice(row)
}

func maxSlice(slice []int) int {
	maxVal := slice[0]
	for _, val := range slice {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
