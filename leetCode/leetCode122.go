package leetCode

import "fmt"

/*
122. 买卖股票的最佳时机 II
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。
*/

func TestLeetCode122() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices1))
}

// 连接递增的数量
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices)) // 表示第i天的最大利润
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(prices)-1]
}
