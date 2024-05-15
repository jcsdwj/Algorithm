package leetCode

import "fmt"

// 289. 生命游戏
/*
根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：
1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。
给你 m x n 网格面板 board 的当前状态，返回下一个状态。
*/

func TestLeetCode289() {
	//board1 := [][]int{
	//	{0, 1, 0},
	//	{0, 0, 1},
	//	{1, 1, 1},
	//	{0, 0, 0},
	//}
	board2 := [][]int{{1}}
	//gameOfLife(board1)
	gameOfLife(board2)
	fmt.Println(board2)
}

func gameOfLife(board [][]int) {
	length := len(board)
	if length == 0 {
		return
	}
	width := len(board[0])

	temp := make([][]int, length)
	for i := 0; i < length; i++ {
		temp[i] = make([]int, width)
	}

	if length == 1 && width == 1 {
		copy(board, temp)
		return
	}

	// 考虑一行一列的情况
	if length == 1 {
		for i := 1; i < width-1; i++ {
			if board[0][i-1] == 1 && board[0][i+1] == 1 {
				temp[0][i] = board[0][i]
			}
		}
		copy(board, temp)
		return
	}

	if width == 1 {
		for i := 1; i < length-1; i++ {
			if board[i-1][0] == 1 && board[i+1][0] == 1 {
				temp[i][0] = board[i][0]
			}
		}
		copy(board, temp)
		return
	}

	// 计算状态
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			count := 0
			if i == 0 {
				// 第1行情况
				if j == 0 {
					// 第1列
					count = board[0][1] + board[1][0] + board[1][1] // 右下对
				} else if j == width-1 {
					// 最后1列
					count = board[0][j-1] + board[1][j] + board[1][j-1] // 左下对
				} else {
					count = board[i][j-1] + board[i][j+1] + board[i+1][j] + board[i+1][j-1] + board[i+1][j+1] // 左右下对
				}

			} else if i == length-1 {
				// 最后一行情况
				if j == 0 {
					// 第1列
					count = board[i-1][j] + board[i][j+1] + board[i-1][j+1]
				} else if j == width-1 {
					// 最后1列
					count = board[i-1][j] + board[i][j-1] + board[i-1][j-1]
				} else {
					count = board[i][j-1] + board[i][j+1] + board[i-1][j-1] + board[i-1][j+1] + board[i-1][j]
				}
			} else {
				if j == 0 {
					// 第1列
					count = board[i-1][j] + board[i+1][j] + board[i][j+1] + board[i-1][j+1] + board[i+1][j+1]
				} else if j == width-1 {
					// 最后1列
					count = board[i-1][j] + board[i+1][j] + board[i][j-1] + board[i-1][j-1] + board[i+1][j-1]
				} else {
					// 中间位置
					count = board[i][j-1] + board[i][j+1] + board[i-1][j] + board[i+1][j] +
						board[i-1][j-1] + board[i-1][j+1] + board[i+1][j-1] + board[i+1][j+1]
					// 前后左右对
				}
			}

			if count > 3 || count < 2 {
				temp[i][j] = 0
			} else if count == 3 && board[i][j] == 0 {
				temp[i][j] = 1
			} else {
				temp[i][j] = board[i][j]
			}
		}
	}
	copy(board, temp)
}
