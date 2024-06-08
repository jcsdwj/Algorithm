/**
  @author: 伍萬
  @since: 2024/6/8
  @desc: //TODO
**/

package leetCode

import "fmt"

// 3040. 相同分数的最大操作数目 II

func TestLeetCode3040() {
	//nums1 := []int{3, 2, 1, 2, 3, 4}
	//nums2 := []int{3, 2, 6, 1, 4}
	nums3 := []int{1, 9, 7, 3, 2, 7, 4, 12, 2, 6}
	//fmt.Println(maxOperations3040(nums1))
	//fmt.Println(maxOperations3040(nums2))
	fmt.Println(maxOperations3040(nums3))
}

// 删除前两个
// 删除后两个
// 删除第一个和最后一个
// 分数最大化

// 递归法超时
func maxOperations3040(nums []int) int {
	// 选择第一个操作后 后面的都是固定的
	v1 := nums[0]
	v2 := nums[1]
	v3 := nums[len(nums)-2]
	v4 := nums[len(nums)-1]
	max1 := getMaxValue(nums[2:], v1+v2)
	max2 := getMaxValue(nums[:len(nums)-2], v3+v4)
	max3 := getMaxValue(nums[1:len(nums)-1], v1+v4)
	return 1 + max(max(max1, max2), max3)
}

func getMaxValue(nums []int, value int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	v1 := nums[0]
	v2 := nums[1]
	v3 := nums[len(nums)-2]
	v4 := nums[len(nums)-1]
	var max1 = 0
	var max2 = 0
	var max3 = 0
	// 第一种情况
	if v1+v2 == value {
		max1 = 1 + getMaxValue(nums[2:], value)
	}
	// 第二种情况
	if v3+v4 == value {
		max2 = 1 + getMaxValue(nums[:len(nums)-2], value)
	}
	// 第三种情况
	if v1+v4 == value {
		max3 = 1 + getMaxValue(nums[1:len(nums)-1], value)
	}

	return max(max(max1, max2), max3)
}

// 记忆化搜索
func maxOperations3040_2(nums []int) int {
	n := len(nums)
	memo := make([][]int, n)

	helper := func(i, j, target int) int {
		for i := range memo {
			memo[i] = make([]int, n)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		var dfs func(int, int) int
		dfs = func(i, j int) int {
			if i >= j {
				return 0
			}
			if memo[i][j] != -1 {
				return memo[i][j]
			}

			ans := 0
			// 前两个
			if nums[i]+nums[i+1] == target {
				ans = max(ans, 1+dfs(i+2, j))
			}
			if nums[j-1]+nums[j] == target {
				ans = max(ans, 1+dfs(i, j-2))
			}
			if nums[i]+nums[j] == target {
				ans = max(ans, 1+dfs(i+1, j-1))
			}
			memo[i][j] = ans // 找到最大
			return ans
		}
		return dfs(i, j)
	}
	res := 0
	res = max(res, helper(0, n-1, nums[0]+nums[n-1]))
	res = max(res, helper(0, n-1, nums[0]+nums[1]))
	res = max(res, helper(0, n-1, nums[n-2]+nums[n-1]))
	return res

}
