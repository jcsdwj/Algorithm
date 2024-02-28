/**
  @author: 伍萬
  @since: 2024/2/26
  @desc: //TODO
**/

package nowcoder

import "math"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 返回最差情况下扔棋子的最小次数
 * @param n int整型 楼层数
 * @param k int整型 棋子数
 * @return int整型
 */
func solve1(n int, k int) int {
	// write code here
	// dp[k,n]表示k个棋子n层楼需要尝试的次数
	// 如果碎了 剩下k-1个棋子 要找的楼层数为i-1 继续尝试dp[k-1,i-10]次
	// 如果没碎 剩下k个棋子 要找的楼层数为n-i 继续尝试dp[k,n-i]次
	dp := [][]int{}
	for i := 0; i < k+1; i++ {
		dp = append(dp, make([]int, n+1))
	}

	// 只有一个棋子
	for i := 1; i < n+1; i++ {
		dp[1][i] = i
	}
	// 只有一层楼
	for i := 1; i < k+1; i++ {
		dp[i][1] = 1
	}

	for i := 2; i < k+1; i++ {
		for j := 2; j < n+1; j++ {
			for m := 1; m <= j; m++ {
				// 从m层开始摔
				dp[i][j] = int(math.Min(float64(dp[i][j]),
					1+math.Max(float64(dp[i][j-m]), float64(dp[i-1][m-1]))))
			}
		}
	}

	return dp[k][n]
}

func solve2(n int, k int) int {
	dp := [][]int{}
	for i := 0; i < k+1; i++ {
		dp = append(dp, make([]int, n+1))
	}

	// 只有一个棋子
	for i := 1; i < n+1; i++ {
		dp[1][i] = i
	}
	// 只有一层楼
	for i := 1; i < k+1; i++ {
		dp[i][1] = 1
	}

	for i := 2; i < k+1; i++ {
		for j := 2; j < n+1; j++ {
			l := 0
			r := j - 1
			for l < r {
				mid := (l + r) / 2
				idx := j - 1 - mid
				if math.Abs(float64(l-r)) == 1 {
					dp[i][j] = 1 + int(math.Max(float64(dp[i][1]), float64(dp[i-1][r])))
					break
				}
				if dp[i-1][mid] == dp[i][idx] {
					dp[i][j] = dp[i][idx] + 1
					break
				}
				if dp[i-1][mid] > dp[i][idx] {
					r = mid
					continue
				}
				if dp[i-1][mid] < dp[i][idx] {
					l = mid
					continue
				}
			}
		}
	}
	return dp[k][n]
}
