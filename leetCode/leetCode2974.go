/**
  @author: 伍萬
  @since: 2024/7/12
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"sort"
)

//2974. 最小数字游戏

func TestLeetCode2974() {
	nums1 := []int{5, 4, 2, 3, 1}
	nums2 := []int{2, 5}
	fmt.Println(numberGame(nums1))
	fmt.Println(numberGame(nums2))
}

func numberGame(nums []int) []int {
	arr := []int{}
	sort.Ints(nums)
	var i = 0
	for ; i < len(nums)-1; i += 2 {
		arr = append(arr, nums[i+1])
		arr = append(arr, nums[i])
	}
	if i == len(nums)-1 {
		// 最后一个数没有加入
		arr = append(arr, nums[i])
	}
	return arr
}
