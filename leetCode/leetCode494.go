/**
  @author: 伍萬
  @since: 2024/6/30
  @desc: //TODO
**/

package leetCode

import "fmt"

// 494. 目标和

func TestLeetCode494() {
	nums1 := []int{1, 1, 1, 1, 1}
	nums2 := []int{1}
	fmt.Println(findTargetSumWays(nums1, 3))
	fmt.Println(findTargetSumWays(nums2, 1))
}

func findTargetSumWays(nums []int, target int) int {
	one := dfs494(nums, target, nums[0], 1)
	two := dfs494(nums, target, -nums[0], 1)
	return one + two
}

func dfs494(nums []int, target int, curVal int, curIndex int) int {
	if curIndex == len(nums) && curVal == target {
		return 1
	}
	if curIndex == len(nums) && curVal != target {
		return 0
	}
	sum := 0
	sum += dfs494(nums, target, curVal+nums[curIndex], curIndex+1) + dfs494(nums, target, curVal-nums[curIndex], curIndex+1)
	return sum
}
