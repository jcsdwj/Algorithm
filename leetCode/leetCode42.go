package leetCode

// 2025-05-15
// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 左右两边不算柱子
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	n := len(height)
	res := 0

	// 数组充当备忘录(左右两边最高的柱子)
	l_max := make([]int, n)
	r_max := make([]int, n)

	l_max[0] = height[0]
	r_max[n-1] = height[n-1]

	// 可以做到连续递增递减
	// 从左向右计算 l_max i点左边最高的柱子
	for i := 1; i < n; i++ {
		l_max[i] = max(height[i], l_max[i-1])
	}
	// 从右向左计算 r_max i点右边最高的柱子
	for i := n - 2; i >= 0; i-- {
		r_max[i] = max(height[i], r_max[i+1])
	}

	// 计算答案
	for i := 1; i < n-1; i++ {
		res += min(l_max[i], r_max[i]) - height[i]
	}
	return res
}
