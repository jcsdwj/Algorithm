/**
  @author: 伍萬
  @since: 2024/5/25
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"sort"
	"strconv"
)

func TestLeetCode386() {
	fmt.Println(lexicalOrder2(10))
}

// 386. 字典序排数
func lexicalOrder(n int) []int {
	strs := []string{}
	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Strings(strs)
	nums := []int{}
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		nums = append(nums, num)
	}
	return nums
}

// 不用sort
func lexicalOrder2(n int) []int {
	nums := make([]int, n)
	num := 1
	for i := range nums {
		nums[i] = num
		if num*10 <= n {
			num *= 10 // 字典序的开始
		} else {
			for num%10 == 9 || num+1 > n { // 字典序的结束
				// 如果到9了 就得开启下一个字典序 如 19->2
				// 如果超过当前n了 也要开启下一个字典序 12->2
				num /= 10
			}
			num++ // 没到就可以继续加1
		}
	}
	return nums
}
