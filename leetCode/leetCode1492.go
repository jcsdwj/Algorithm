/**
  @author: 伍萬
  @since: 2024/5/19
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"math"
)

// 1492. n 的第 k 个因子

//给你两个正整数 n 和 k 。
//
//如果正整数 i 满足 n % i == 0 ，那么我们就说正整数 i 是整数 n 的因子。
//
//考虑整数 n 的所有因子，将它们 升序排列 。请你返回第 k 个因子。如果 n 的因子数少于 k ，请你返回 -1 。

func TestLeetCode1492() {
	fmt.Println(kthFactor(12, 3))
	fmt.Println(kthFactor(7, 2))
	fmt.Println(kthFactor(4, 4))
}

func kthFactor(n int, k int) int {
	// 存在2*2 只能存一个
	nums := []int{1, n}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			var temp []int
			if i != n/i {
				temp = []int{i, n / i}
			} else {
				temp = []int{i}
			}
			nums = append(nums[:len(nums)/2], append(temp, nums[len(nums)/2:]...)...)
		}
	}
	if len(nums) >= k {
		return nums[k-1]
	}
	return -1
}
