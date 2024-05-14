package leetCode

import "fmt"

/*
238. 除自身以外数组的乘积
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/

func TestLeetCode238() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums1))
	fmt.Println(productExceptSelf(nums2))
}

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	dpFront := make([]int, len(nums)) // 表示前i个数的积
	dpRear := make([]int, len(nums))  // 表示后i个数的积

	dpFront[0] = nums[0]
	dpRear[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		dpFront[i] = dpFront[i-1] * nums[i]
	}
	for j := len(nums) - 2; j >= 0; j-- {
		dpRear[j] = dpRear[j+1] * nums[j]
	}

	ans := make([]int, len(nums))
	ans[0] = dpRear[1]
	ans[len(nums)-1] = dpFront[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		// 前i个数乘以i+1后个数
		ans[i] = dpFront[i-1] * dpRear[i+1]
	}
	return ans
}
