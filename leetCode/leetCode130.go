/**
  @author: 伍萬
  @since: 2024/5/26
  @desc: //TODO
**/

package leetCode

// 130. 被围绕的区域

func solve(board [][]byte) {
	// 找不到出路转为X
	if len(board) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])

	// 边界上的 O不会变化 所以与边界O相连的都不会变 其他都是X
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {
				// 从边界O访问
				dfs130(board, i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs130(board [][]byte, i, j int) {
	// 遇到边界返回
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '*' {
		return
	}

	board[i][j] = '*'     // 表示与边界O相连
	dfs130(board, i-1, j) // 上
	dfs130(board, i+1, j) // 下
	dfs130(board, i, j-1) // 左
	dfs130(board, i, j+1) // 右
}
