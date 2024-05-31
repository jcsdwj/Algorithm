package leetCode

// 2965. 找出缺失和重复的数字

// 暴力法
func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	if n == 0 {
		return nil
	}

	// 第一个数存重复 第二个数存缺失
	nums := []int{}
	dict := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dict[grid[i][j]] == true {
				nums = append(nums, grid[i][j])
			}
			dict[grid[i][j]] = true
		}
	}
	for i := 1; i <= n*n; i++ {
		if dict[i] == false {
			nums = append(nums, i)
			break
		}
	}

	return nums
}

// sum-(1+n)*n/2=a-b
// 所有数的平方和减去 公式的 = a^2-b^2
// 数学法
func findMissingAndRepeatedValues2(grid [][]int) []int {
	n := len(grid)
	m := n * n
	d1 := -m * (m + 1) / 2
	d2 := -m * (m + 1) * (m*2 + 1) / 6
	for _, row := range grid {
		for _, x := range row {
			d1 += x
			d2 += x * x
		}
	}
	return []int{(d2/d1 + d1) / 2, (d2/d1 - d1) / 2}
}
