/**
  @author: 伍萬
  @since: 2024/6/11
  @desc: //TODO
**/

package leetCode

import "fmt"

// 419. 甲板上的战舰
// 1*k或k*1

func TestLeetCode419() {
	board1 := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	board2 := [][]byte{}
	fmt.Println(countBattleships(board1))
	fmt.Println(countBattleships(board2))
}

func countBattleships(board [][]byte) int {
	if board == nil {
		return 0
	}
	m := len(board)
	if m == 0 {
		return 0
	}
	n := len(board[0])
	if n == 0 {
		return 0
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				count++
				dfs419(board, i, j)
			}
		}
	}
	return count
}

func dfs419(board [][]byte, x, y int) {
	m := len(board)
	n := len(board[0])

	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}

	if board[x][y] == '.' {
		return
	}
	// 变成.
	board[x][y] = '.'
	// 只能沿着一个方向
	dfs419(board, x-1, y)
	dfs419(board, x+1, y)
	dfs419(board, x, y-1)
	dfs419(board, x, y+1)
}
