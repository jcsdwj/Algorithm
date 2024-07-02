/**
  @author: 伍萬
  @since: 2024/7/2
  @desc: //TODO
**/

package leetCode

import "fmt"

func TestLeetCode3115() {
	fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 3}))
	fmt.Println(maximumPrimeDifference([]int{4, 8, 2, 8}))
}

// 3115. 质数的最大距离

func maximumPrimeDifference(nums []int) int {
	// 从左右各找一个质数
	var start, end = -1, len(nums)
	for i := 0; i < len(nums); i++ {
		if isPrime(nums[i]) {
			start = i
			break
		}
	}
	for j := len(nums) - 1; j >= 0; j-- {
		if isPrime(nums[j]) {
			end = j
			break
		}
	}
	if start == -1 {
		return 0
	}
	return end - start
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
