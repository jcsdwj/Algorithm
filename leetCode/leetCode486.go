package leetCode

// 486. 预测赢家

// 只能取头和尾元素

// 递归法
func predictTheWinner(nums []int) bool {
	return total(nums, 0, len(nums)-1, 1) >= 0
}

func total(nums []int, start, end, turn int) int {
	if start == end {
		return nums[start] * turn
	}

	//
	scoreStart := nums[start]*turn + total(nums, start+1, end, -turn) // 选头部的情况
	scoreEnd := nums[end]*turn + total(nums, start, end-1, -turn)     // 选尾部的情况
	return max(scoreStart*turn, scoreEnd*turn) * turn
}

// 动态规划法
func predictTheWinner2(nums []int) bool {
	return false
}
