/**
  @author: 伍萬
  @since: 2024/6/7
  @desc: //TODO
**/

package leetCode

import "fmt"

// 3038. 相同分数的最大操作数目 I

func TestLeetCode3038() {
	nums1 := []int{3, 2, 1, 4, 5}
	nums2 := []int{3, 2, 6, 1, 4}
	fmt.Println(maxOperations(nums1))
	fmt.Println(maxOperations(nums2))
}

func maxOperations(nums []int) int {
	sum := nums[0] + nums[1]

	count := 1
	for i := 2; i < len(nums)-1; i += 2 {
		if nums[i]+nums[i+1] == sum {
			count++
		} else {
			break
		}
	}
	return count
}
