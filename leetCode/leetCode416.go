package leetCode

import "math"

// 416. 分割等和子集

// 给你一个只包含正整数的非空数组nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	var max = math.MinInt
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if sum%2 != 0 {
		return false
	}

	// 子集之和为sum/2
	sum = sum / 2

	// 最大值大于sum的一半
	if sum < max {
		return false
	}

	dp := make([][]bool, n) // 表示前i个数能否凑到j的值
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	// 不可能大于
	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j < sum+1; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				// 不取这个值
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][sum] // 返回是否有这个值
}
