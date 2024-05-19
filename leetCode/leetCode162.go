/**
  @author: 伍萬
  @since: 2024/5/19
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"math"
)

func TestLeetCode162() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums1))
	fmt.Println(findPeakElement(nums2))
}

// 162. 寻找峰值 第一个数和最左边的数大于旁边的数即可
// log n
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return 0
}

// 二分查找的方法  只要有一个峰值就可以
func findPeakElement2(nums []int) int {
	n := len(nums)
	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			// 说明大值在右边 可以抛弃左边（求出一个就好，下面同理）
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
