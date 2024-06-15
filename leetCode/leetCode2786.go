/**
  @author: 伍萬
  @since: 2024/6/14
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"math"
)

//2786. 访问数组中的位置使分数最大

func TestLeetCode2786() {
	nums1 := []int{2, 3, 6, 1, 9, 2}
	nums2 := []int{2, 4, 6, 8}
	fmt.Println(maxScore(nums1, 5))
	fmt.Println(maxScore(nums2, 3))
}

// 奇偶性只能变一次
// 若与第一个数不一样 则后面也必须不一样
// 知道最后移动元素为奇数的情况，最后移动元素为偶数的情况
// 每次变换都会扣除x 当前写法忽略了
// 最后为奇数和最后为偶数的最大值情况
func maxScore(nums []int, x int) int64 {
	res := int64(nums[0])
	dp := [2]int64{math.MinInt32, math.MinInt32}
	dp[nums[0]%2] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		flag := nums[i] % 2

		// 以当前值结尾的情况
		cur := maxInt64(dp[flag]+int64(nums[i]), dp[1-flag]-int64(x)+int64(nums[i]))
		res = maxInt64(res, cur)
		dp[flag] = maxInt64(dp[flag], cur)
	}
	return res
}

func maxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
