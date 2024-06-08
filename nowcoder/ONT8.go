/**
  @author: 伍萬
  @since: 2024/6/7
  @desc: //TODO
**/

package nowcoder

import (
	"Algorithm/common"
	"math"
)

// ONT8 牛牛的冒险旅程

// 判断是否有环，有返回-1 没有求最大公约数
func gcdInCycle(head *common.ListNode) int {
	// write code here
	dict := make(map[int]int)
	l := head

	nums := []int{}
	minVal := math.MaxInt
	count := 0
	for l != nil {
		dict[l.Val]++
		if dict[l.Val] == 1 {
			nums = append(nums, l.Val)
		}
		count++
		if l.Val < minVal {
			minVal = l.Val
		}
		l = l.Next
	}

	// 无环情况
	if count == len(dict) {
		return -1
	}

	return getMostCommonDivisor(nums, minVal)
}

func getMostCommonDivisor(nums []int, minValue int) int {
	i := minValue
	for ; i >= 1; i-- {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]%i != 0 {
				continue
			}
			count++
		}
		if count == len(nums) {
			break
		}
	}
	return i
}
