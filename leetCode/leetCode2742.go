/**
  @author: 伍萬
  @since: 2024/6/28
  @desc: //TODO
**/

package leetCode

import "math"

// 2742. 给墙壁刷油漆

// 没理解

// 最小开销  免费的效率是1
// 花钱最少 耗时最多的

// 花费cost[i]可以完成 1+time[i]个任务
// 至少有个任务是付费完成
// time至少为1 也就是至少n/2被免费完成
// 从最小的n/2个cost里挑
// 本质上是统计哪些cost的组合的time之和大于整个数组长度
// 状态转移方程考虑第i位油漆工是付费还是免费工作
// 刷第n-1道墙 在time[n-1]时间里
// 不刷n-1道墙

// f(i,j)表示前i位油漆匠 免费工作次数为j的最小开销

// 动态规划法
func paintWalls1(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, 2*n+1) // 最坏情况都要花钱 最好情况都不用花钱 免费次数j
	for k := range f {
		f[k] = math.MaxInt / 2
	}
	f[n] = 0

	for i := 0; i < n; i++ {
		g := make([]int, n*2+1)
		for k := range g {
			g[k] = math.MaxInt / 2
		}
		for j := 0; j <= n*2; j++ {
			// 付费
			g[min(j+time[i], n*2)] = min(g[min(j+time[i], n*2)], f[j]+cost[i])
			// 免费
			if j > 0 {
				g[j-1] = min(g[j-1], f[j])
			}
		}
		f = g
	}
	res := math.MaxInt
	for i := n; i < len(f); i++ {
		res = min(res, f[i])
	}
	return res
}

// 空间优化
func paintWalls2(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, n+2)
	for k := range f {
		f[k] = math.MaxInt / 2
	}
	f[0] = 0
	for i := 0; i < n; i++ {
		for j := n + 1; j >= 0; j-- {
			f[min(j+time[i], n)+1] = min(f[min(j+time[i], n)+1], f[j]+cost[i])
		}
	}
	return min(f[n], f[n+1])
}

// 付费与免费的时间差
func paintWalls3(cost []int, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	var dfs func(int, int) int // 表示刷前i堵墙 付费时间减去免费时间之和为j 可以免费的时间
	dfs = func(i, j int) int {
		if j > i { // 剩余的墙都可以免费刷
			return 0
		}
		if i < 0 { // 上面 if 不成立，意味着 j < 0，不符合题目要求
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j+n] // 加上偏移量 n，防止出现负数
		if *p != -1 {      // 之前计算过
			return *p
		}
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, 0)
}

// 付费刷墙个数 + 免费刷墙个数 =n
// 付费刷墙时间之和必须 ≥ 免费刷墙个数
// 转成背包问题
func paintWalls4(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1 // 注意这里加一了
		for j := n; j > 0; j-- {
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
	}
	return f[n]
}
