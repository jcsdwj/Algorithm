package leetCode

import (
	"fmt"
	"strconv"
)

// 60. 排列序列
// 1..n的第k种排列方式

// 计算总共有多少种排列 每个数有n-1阶乘种方案 总数除以k 等于几就是第几个数开头
// 计算是哪个数开头的
// 去掉这个数 计算是这些数的多少种排列 减去

func TestLeetCode60() {
	fmt.Println(getPermutation(2, 1))
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
	fmt.Println(getPermutation(3, 2))
}

func getPermutation(n int, k int) string {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i) // 构建数组
	}
	return getStr(nums, k, n)
}

// 每次从剩余数取值
func getStr(nums []int, k int, length int) string {
	// 只剩一个数直接取
	if length == 1 {
		return strconv.Itoa(nums[0])
	}

	turn := getWays(length - 1) // 一个数开头有多少种方案 3

	// 算第一个数在第几个位置(数组的位置)  123 132 213 231 312 321 3
	number := (k - 1) / turn

	// 拿这个字符 递归操作 数组少了一个数 k为减去已经经过了多少轮 这就是剩余的次数 数组长度少1
	return strconv.Itoa(nums[number]) + getStr(append(nums[:number], nums[number+1:]...), k-turn*number, length-1)
}

// 获取总方案数 也就是阶乘
func getWays(k int) int {
	sum := 1
	for i := k; i >= 1; i-- {
		sum *= i
	}
	return sum
}
