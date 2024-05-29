package leetCode

// 312. 戳气球

// 动态规划法
func maxCoins(nums []int) int {
	// 添气球达到nums
	// dp[i][j] =dp[i][k]+dp[k][j]+points[i]*points[j]*points[k]
	// 戳破i到j之间的所有气球得到的最高分数

	n := len(nums)
	points := make([]int, n+2) // 添加两侧的气球
	points[0] = 1
	points[n+1] = 1

	for i := 1; i <= n; i++ {
		points[i] = nums[i-1]
	}

	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}

	// 从左往右
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
}

// 记忆化搜索

func maxCoins2(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	rec := make([][]int, n+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, n+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	return solve312(0, n+1, val, rec)
}

// 填满left到right能够得到的金币数
func solve312(left, right int, val []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rec[left][right] != -1 {
		return rec[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		sum += solve312(left, i, val, rec) + solve312(i, right, val, rec)
		rec[left][right] = max(rec[left][right], sum)
	}
	return rec[left][right]
}
