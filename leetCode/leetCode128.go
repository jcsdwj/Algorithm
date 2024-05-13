package leetCode

import (
	"fmt"
)

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

func TestLeetCode128() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums3 := []int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}
	nums4 := []int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999}
	fmt.Println(longestConsecutive(nums1))
	fmt.Println(longestConsecutive(nums2))
	fmt.Println(longestConsecutive(nums3))
	fmt.Println(longestConsecutive(nums4))
}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0
	for num := range numSet {
		// 我们只对序列中的第一个数字进行计算，即其前一个数字不存在的数字(如果前一个数不存在再进行)
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// 计算当前数字所在的连续序列长度
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// 更新最长连续序列长度
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
