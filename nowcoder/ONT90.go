/**
  @author: 伍萬
  @since: 2024/5/9
  @desc: //TODO
**/

package nowcoder

import (
	"fmt"
	"sort"
	"strconv"
)

func SolveONT90() {
	fmt.Println(orderArray(100))
}

// 字典序排列
func orderArray(n int) []int {
	// write code here
	strs := []string{}
	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Strings(strs)
	nums := []int{}
	for i := 0; i < len(strs); i++ {
		v, _ := strconv.Atoi(strs[i])
		nums = append(nums, v)
	}
	return nums
}

// 方法2
var ans []int

func orderArray2(n int) []int {
	// write code here
	for i := 1; i <= 9; i++ {
		dfs(i, n)
	}
	return ans
}

func dfs(num, n int) {
	if num > n {
		return
	}
	ans = append(ans, num)
	for i := 0; i <= 9; i++ {
		dfs(num*10+i, n)
	}
}
