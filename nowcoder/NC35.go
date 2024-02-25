/**
  @author: 伍萬
  @since: 2024/2/24
  @desc: //TODO
**/

package nowcoder

import (
	"fmt"
	"math"
)

func TestNC35() {
	//fmt.Println(minEditCost("abc", "adc", 5, 3, 2))
	fmt.Println(minEditCost("abc", "adc", 5, 3, 100))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * min edit cost
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @param ic int整型 insert cost
 * @param dc int整型 delete cost
 * @param rc int整型 replace cost
 * @return int整型
 */
func minEditCost(str1 string, str2 string, ic int, dc int, rc int) int {
	// write code here
	l1 := len(str1)
	l2 := len(str2)
	dp := [][]int{}
	for i := 0; i < l1+1; i++ {
		dp = append(dp, make([]int, l2+1))
	}
	// dp[i][j]表示str1前i个字符到str2前j个字符的代价 str1=>str2
	dp[0][0] = 0
	dp[0][1] = ic
	dp[1][0] = dc

	for i := 1; i < l1+1; i++ {
		dp[i][0] = dc * i
	}

	for j := 1; j < l2+1; j++ {
		dp[0][j] = ic * j
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			var a, b, c int
			if str1[i-1] == str2[j-1] {
				// 字符相等情况
				a = dp[i-1][j-1]
			} else {
				a = dp[i-1][j-1] + int(math.Min(float64(rc), float64(ic+dc)))
			}
			b = dp[i][j-1] + ic
			c = dp[i-1][j] + dc
			dp[i][j] = int(math.Min(float64(a), math.Min(float64(b), float64(c))))
		}
	}

	return dp[l1][l2]
}
