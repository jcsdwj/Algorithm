package leetCode

// 718. 最长重复子数组
// 给两个整数数组 nums1 和 nums2 ，返回两个数组中公共的 、长度最长的子数组的长度 。

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1) // 表示第i个元素和第j个元素的最长公共子数组长度
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}
	return maxLen
}
