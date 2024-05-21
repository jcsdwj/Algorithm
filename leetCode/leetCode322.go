package leetCode

import (
	"fmt"
	"math"
)

// 322. 零钱兑换

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

// 你可以认为每种硬币的数量是无限的。

func TestLeetCode322() {
	coins1 := []int{1, 2, 5}
	coins2 := []int{2}
	coins3 := []int{1}
	fmt.Println(coinChange(coins1, 11))
	fmt.Println(coinChange(coins2, 3))
	fmt.Println(coinChange(coins3, 0))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i]表示i钱需要多少硬币
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		if coins[i] < amount+1 {
			dp[coins[i]] = 1 // 需要一枚硬币
		}
	}

	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				// 超额了
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
