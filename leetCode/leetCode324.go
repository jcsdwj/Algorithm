package leetCode

import (
	"fmt"
	"sort"
)

// 324. 摆动排序 II

// 给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
// 1,3,2,2,3,1  2,3,1,3,1,2

func TestLeetCode324() {
	nums1 := []int{1, 3, 2, 2, 3, 1} // 1 2 1 3 2 3
	nums2 := []int{1, 5, 1, 1, 6, 4}
	nums3 := []int{1, 1, 2, 1, 2, 2, 1}
	nums4 := []int{4, 5, 5, 6} // 5645
	wiggleSort(nums1)
	wiggleSort(nums2)
	wiggleSort(nums3)
	wiggleSort(nums4)
}

// 每次都从数组最大的开始选
// 分成两部分 右边数组的最大值肯定大于左边的最大值
func wiggleSort(nums []int) {
	n := len(nums)
	if n == 1 || n == 0 {
		return
	}
	sort.Ints(nums)

	left := []int{}
	right := []int{}
	for i := 0; i < n; i++ {
		if i < (n+1)/2 {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			nums[i] = left[len(left)-1]
			left = left[:len(left)-1]
		} else {
			nums[i] = right[len(right)-1]
			right = right[:len(right)-1]
		}
	}

	fmt.Println(nums)
}
