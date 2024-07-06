/**
  @author: 伍萬
  @since: 2024/7/6
  @desc: //TODO
**/

package leetCode

import "fmt"

func TestLeetCode3101() {
	//nums1 := []int{0, 1, 1, 1}
	nums2 := []int{1, 0, 1, 0}
	nums3 := []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1}
	//fmt.Println(countAlternatingSubarrays2(nums1))
	fmt.Println(countAlternatingSubarrays2(nums2))
	fmt.Println(countAlternatingSubarrays2(nums3))
}

// 3101. 交替子数组计数
// 不存在两个相邻元素值相同
// 超时
func countAlternatingSubarrays(nums []int) int64 {
	// 从单个开始统计
	sum := int64(len(nums))
	for i := 2; i <= len(nums); i++ {
		// 子数组的长度
		for j := 0; j+i <= len(nums); j++ {
			sum += isSubarrays(nums[j : j+i])
		}
	}
	return sum
}

func isSubarrays(nums []int) int64 {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return 0
		}
	}
	return 1
}

func countAlternatingSubarrays2(nums []int) int64 {
	if len(nums) == 1 {
		return 1
	}
	// 新增一个元素 如果与上一个相同 只会加1 不同则加1+2*(上一轮)
	dp := make([]int, len(nums))
	dp[0] = 1
	if nums[0] == nums[1] {
		dp[1] = 2
	} else {
		dp[1] = 3
	}
	// 当前值为作为i的个数

	for i := 2; i < len(dp); i++ {
		if nums[i] == nums[i-1] {
			// 相同情况 只会+1
			dp[i] = 1 + dp[i-1]
		} else {
			// 不同情况 上一个减去上上一个的增加的个数 加单独构成数组的1
			dp[i] = 1 + 2*dp[i-1] - dp[i-2]
		}
	}
	return int64(dp[len(dp)-1])
}

func countAlternatingSubarrays3(nums []int) int64 {
	var ans int64
	cnt := 0 // 统计新生成的交替子数组
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			cnt = 1
		} else {
			cnt++
		}
		ans += int64(cnt) // 有 cnt 个以 i 为右端点的交替子数组
	}
	return ans
}
