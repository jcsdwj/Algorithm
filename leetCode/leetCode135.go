package leetCode

import "fmt"

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目

*/

func TestLeetCode135() {
	// ratings1 := []int{1, 0, 2}
	// ratings2 := []int{1, 2, 2}
	ratings3 := []int{1, 2, 87, 87, 87, 2, 1}
	// fmt.Println(candy(ratings1))
	// fmt.Println(candy(ratings2))
	fmt.Println(candy(ratings3))
}

func candy(ratings []int) int {
	// 基础n个
	// 连续大于或连续小于的次数
	// 从左往右看
	// 既比左边大又比右边大的会算重复
	// 求相对高度
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	} else if len(ratings) == 2 {
		if ratings[0] == ratings[1] {
			return 2
		} else {
			return 3
		}
	}
	ans := 0

	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++ // 左边大于右边的次数
		} else {
			right = 1
		}
		ans += max(left[i], right) // 二者最大高度?
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
