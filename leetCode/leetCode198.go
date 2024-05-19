/**
  @author: 伍萬
  @since: 2024/5/19
  @desc: //TODO
**/

package leetCode

import "fmt"

//198. 打家劫舍

func TestLeetCode198() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums1))
	fmt.Println(rob(nums2))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)) // 如果偷i的最多钱
	dp[0] = nums[0]
	dp[1] = nums[1]

	if len(nums) == 2 {
		if dp[0] > dp[1] {
			return dp[0]
		}
		return dp[1]
	}

	dp[2] = nums[2] + nums[0]
	// 可以连跨两个不能跨三个
	for i := 3; i < len(nums); i++ {
		if dp[i-2] > dp[i-3] {
			dp[i] = nums[i] + dp[i-2]
		} else {
			dp[i] = nums[i] + dp[i-3]
		}
	}

	n := len(nums)
	if dp[n-1] > dp[n-2] {
		return dp[n-1]
	}
	return dp[n-2]
}

// 另一种方法
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)) // 当前最大利润
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
