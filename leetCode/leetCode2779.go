/**
  @author: 伍萬
  @since: 2024/6/15
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"sort"
)

// 2779. 数组的最大美丽值

func TestLeetCode2779() {
	//nums1 := []int{4, 6, 1, 2}
	//nums2 := []int{1, 1, 1, 1}
	//nums3 := []int{49, 26}
	nums4 := []int{43, 86, 33, 18}
	//fmt.Println(maximumBeauty(nums1, 2))
	//fmt.Println(maximumBeauty(nums2, 2))
	//fmt.Println(maximumBeauty(nums3, 12)) // 可以49-12 26+11
	fmt.Println(maximumBeauty(nums4, 23))
}

// 每个数往后找交集的最大值
// 保证不了跨过第二个直接与第三个数交集
func maximumBeauty(nums []int, k int) int {
	// 不同数的最大交集
	temp := [][]int{}
	for i := 0; i < len(nums); i++ {
		temp = append(temp, []int{nums[i] - k, nums[i] + k})
	}

	var maxCount int
	for i := 0; i < len(temp); i++ {
		count := 1
		interval := []int{temp[i][0], temp[i][1]}
		for j := i + 1; j < len(temp); j++ {
			if temp[j][0] > interval[1] || temp[j][1] < interval[0] {
				// 没有交集的直接下一个
				continue
			}

			// 有交集情况 更新区间
			if temp[j][0] > interval[0] {
				interval[0] = temp[j][0]
			}
			if temp[j][1] < interval[1] {
				interval[1] = temp[j][1]
			}
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

// 排序后再取交集
func maximumBeauty2(nums []int, k int) int {
	res, n := 0, len(nums)
	sort.Ints(nums)

	// i作为右节点，找到离自己最远的j的位置 求出一个结果
	// i和j都往右移动
	for i, j := 0, 0; i < n; i++ {
		for nums[i]-2*k > nums[j] {
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}
