/**
  @author: 伍萬
  @since: 2024/6/24
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
)

// 503. 下一个更大元素 II
func TestLeetCode503() {
	nums1 := []int{1, 2, 1}
	nums2 := []int{1, 2, 3, 4, 3}
	nums3 := []int{5, 4, 3, 2, 1}
	fmt.Println(nextGreaterElements(nums1))
	fmt.Println(nextGreaterElements(nums2))
	fmt.Println(nextGreaterElements(nums3))
}

// 循环搜索下一个比自己更大的值
func nextGreaterElements(nums []int) []int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		for j := (i + 1) % len(nums); j != i; j = (j + 1) % len(nums) {
			if nums[i] < nums[j] {
				// 找到值了
				dp[i] = nums[j]
				j = i - 1
			}
		}
	}

	return dp
}

// 使用单调栈方法会更快 O(n)
