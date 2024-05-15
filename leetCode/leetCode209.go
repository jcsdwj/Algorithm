/**
  @author: 伍萬
  @since: 2024/5/15
  @desc: //TODO
**/

package leetCode

import "math"

// 209. 长度最小的子数组

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < length; i++ {
		if nums[i] > target {
			return 1
		}
		sum += nums[i]
	}
	if sum < target {
		return 0
	}

	ans := math.MaxInt
	// 滑动窗口
	start, end := 0, 0
	sum = 0
	for end < length {
		sum += nums[end]
		for sum >= target {
			// 依然大于则移动头部
			ans = min(ans, end-start+1)
			sum -= nums[end]
			start++
		}

		// 直到小于则移动尾部
		end++
	}

	return ans
}
