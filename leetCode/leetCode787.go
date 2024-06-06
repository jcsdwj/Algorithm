package leetCode

import (
	"fmt"
	"math"
)

// 787. K 站中转内最便宜的航班

// 超时
func TestLeetCode787() {
	edges := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	edges2 := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}

	fmt.Println(findCheapestPrice2(3, edges, 0, 2, 1))
	fmt.Println(findCheapestPrice2(3, edges, 0, 2, 0))
	fmt.Println(findCheapestPrice2(3, edges2, 0, 2, 0))
}

// n个点 flights票价 src起点 dst终点 k中转次数(最多经过k中转)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	path := make([][]int, n)
	for i := 0; i < n; i++ {
		path[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = -1
		}
	}
	for i := 0; i < len(flights); i++ {
		start := flights[i][0]
		end := flights[i][1]
		weight := flights[i][2]
		path[start][end] = weight
	}

	return findLowest(n, path, src, dst, k)
}

// src最多经过k到dst的最短路径
func findLowest(n int, path [][]int, src, dst, k int) int {
	if k == 0 {
		return path[src][dst]
	}
	var minPath = path[src][dst]
	for i := 0; i < n; i++ {
		if path[src][i] == -1 { // 没有路径到中间节点略过
			continue
		}
		curPath := findLowest(n, path, i, dst, k-1)
		if curPath == -1 { // 没有路径直接略过
			continue
		}

		// 新路径更小 或者 找到了路径
		if minPath > curPath+path[src][i] || minPath == -1 {
			minPath = curPath + path[src][i]
		}
	}

	return minPath
}

// 采用动态规划
func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	// dp[x][y] 经过x个站点到达y点
	// dp[x][y]=dp[x-1][y], dp[x-1][z]+cost[z][y]
	// 少一次达到情况和一个中间位置到达情况
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < k+1; i++ {
		for j := 0; j < n; j++ {
			// 先设为-1 都没有位置能到
			dp[i][j] = -1
		}
	}

	for i := 0; i < len(flights); i++ {
		if flights[i][0] == src {
			// src直接能到的点
			dp[0][flights[i][1]] = flights[i][2]
		}
	}

	for i := 1; i <= k; i++ {
		for _, flight := range flights {
			start := flight[0]
			end := flight[1]
			weight := flight[2]

			// 少一步达到end的情况
			firstVal := dp[i-1][end]

			// 少一步达到start，再从start到end
			secondVal := dp[i-1][start]

			// 说明上一步没有直接到end位置的
			if firstVal == -1 && secondVal != -1 {
				if dp[i][end] == -1 {
					dp[i][end] = secondVal + weight
				} else {
					dp[i][end] = min(dp[i][end], secondVal+weight)
				}
			}

			// 上一步既不能到end也不能到start
			if firstVal == -1 && secondVal == -1 {
				continue
			}

			// 不能到end能到start
			if firstVal != -1 && secondVal == -1 {
				if dp[i][end] == -1 {
					dp[i][end] = firstVal
				} else {
					dp[i][end] = min(dp[i][end], firstVal)
				}
			}

			// 都能到
			if firstVal != -1 && secondVal != -1 {
				if dp[i][end] == -1 {
					dp[i][end] = min(firstVal, secondVal+weight)
				} else {
					dp[i][end] = min(min(firstVal, secondVal+weight), dp[i][end])
				}
			}
		}
	}

	var minWeight = math.MaxInt
	for i := 0; i <= k; i++ {
		if dp[i][dst] < minWeight {
			// 不能到的为-1 肯定最小
			minWeight = dp[i][dst]
		}
	}
	return minWeight
}

func findCheapestPrice3(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1 // 最大值情况
	f := make([][]int, k+2)

	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}

	// dp[i][j]表示经过i个点到j的最小路径值

	f[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			f[t][i] = min(f[t][i], f[t-1][j]+cost)
			// 为什么不是f[t-1][i]呢
		}
	}

	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, f[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}
