/**
  @author: 伍萬
  @since: 2024/7/10
  @desc: //TODO
**/

package leetCode

import "fmt"

//2970. 统计移除递增子数组的数目 I

func TestLeetCode2970() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{6, 5, 7, 8}
	nums3 := []int{8, 7, 6, 6}
	fmt.Println(incremovableSubarrayCount(nums1))
	fmt.Println(incremovableSubarrayCount(nums2))
	fmt.Println(incremovableSubarrayCount(nums3))
}

// 移除一个元素剩余元素严格递增的个数
// 如果一个数组递增，那么移除它所有的元素都递增
// 假设数组为n 则结果为n的递归
func incremovableSubarrayCount(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isIncreasing(nums, i, j) {
				res++
			}
		}
	}
	return res
}

func isIncreasing(nums []int, l, r int) bool {
	for i := 1; i < len(nums); i++ {
		if i >= l && i <= r+1 {
			continue
		}
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	if l-1 >= 0 && r+1 < len(nums) && nums[r+1] <= nums[l-1] {
		return false
	}
	return true
}
