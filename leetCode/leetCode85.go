package leetCode

// 85. 最大矩形

// 暴力法
func maximalRectangle(matrix [][]byte) int {
	var ans int
	if len(matrix) == 0 {
		return ans
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m) // left[i][j]表示 matrix[i][j]结尾 i行连续1的个数

	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := left[i][j] // 当前最长宽度
			area := width

			// 往上找
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])  // 比这个宽度小的
				area = max(area, (i-k+1)*width) // 求面积
			}
			ans = max(ans, area)
		}
	}
	return ans
}

// 单调栈
func maximalRectangle2(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = max(ans, area)
		}
	}
	return
}
