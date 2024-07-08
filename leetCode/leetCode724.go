/**
  @author: 伍萬
  @since: 2024/7/8
  @desc: //TODO
**/

package leetCode

import "fmt"

func TestLeetCode724() {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	nums2 := []int{1, 2, 3}
	nums3 := []int{2, 1, -1}
	fmt.Println(pivotIndex(nums1))
	fmt.Println(pivotIndex(nums2))
	fmt.Println(pivotIndex(nums3))
}

// 724. 寻找数组的中心下标
// 存在多个返回最左边 不存在返回-1
func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	var temp int // 用来保存左边值之和 初始为0

	// 第一个数作为中心的情况
	if temp == sum-nums[0] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		temp = temp + nums[i-1]
		if temp == sum-temp-nums[i] {
			return i
		}
	}
	return -1
}
