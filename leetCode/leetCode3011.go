/**
  @author: 伍萬
  @since: 2024/7/13
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"math/bits"
	"sort"
)

//3011. 判断一个数组是否可以变为有序

func TestLeetCode3011() {
	nums1 := []int{8, 4, 2, 30, 15}
	nums2 := []int{1, 2, 3, 4, 5}
	nums3 := []int{3, 16, 8, 4, 2}
	fmt.Println(canSortArray(nums1))
	fmt.Println(canSortArray(nums2))
	fmt.Println(canSortArray(nums3))
}

type numToBit struct {
	num int
	bit int
}

type numToBits []numToBit

func (p numToBits) Len() int { return len(p) }
func (p numToBits) Less(i, j int) bool {
	return p[i].bit < p[j].bit
}
func (p numToBits) Swap(i, j int) {
	if p[i].bit == p[j].bit {
		p[i], p[j] = p[j], p[i]
	}
}

// 3011. 判断一个数组是否可以变为有序
// 二进制1的个数相同的可以换位置
// 位数相同且连续的可以换位置保持有序
func canSortArray(nums []int) bool {
	countBits := []int{}
	for i := 0; i < len(nums); i++ {
		countBits = append(countBits, countBit(nums[i]))
	}

	countBits = append(countBits, -1)

	start, end := 0, 0
	for end <= len(nums) {
		if countBits[end] == countBits[start] {
			end++
			continue
		}
		// 不等于的时候
		sort.Ints(nums[start:end])
		start = end
		end++
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}

func countBit(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	sum := 0
	if num%2 == 0 {
		sum += countBit(num / 2)
	} else {
		sum += 1 + countBit(num/2)
	}
	return sum
}

func canSortArray2(nums []int) bool {
	lastCnt := 0
	lastGroupMax := 0
	curGroupMax := 0

	for _, num := range nums {
		curCnt := bits.OnesCount(uint(num)) // 当前数的1的个数
		if curCnt == lastCnt {
			curGroupMax = max(curGroupMax, num) // 保留当前组最大值
		} else {
			// 向后走
			lastCnt = curCnt
			lastGroupMax = curGroupMax
			curGroupMax = num
		}

		// 上一组连续的最大值
		if num < lastGroupMax {
			return false
		}
	}
	return true
}
