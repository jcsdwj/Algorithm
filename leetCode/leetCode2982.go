/**
  @author: 伍萬
  @since: 2024/5/30
  @desc: //TODO
**/

package leetCode

import "sort"

// 2982. 找出出现至少三次的最长特殊子字符串 II
// 一定是连续的 例如aaa bb c
// 找到每种字符串最长的三种长度
// 没有的话肯定是-1
// 假设字符c长度为l0 l1 l2
// 取值方案有
// 从最长的l0取 l0-2作为子串
// 假设l0=l1 取l0-1作为子串
// 假设l0大于l1 取三个l1作为子串
// 从最长 次长 第三长取l2(不可以直接取次长吗)
func maximumLength(s string) int {
	n := len(s)
	cnt := make([][]int, 26)
	count := 0
	for i := range s {
		count++
		if i+1 == n || s[i] != s[i+1] {
			// 统计连续长度
			cnt[s[i]-'a'] = append(cnt[s[i]-'a'], count)
			count = 0
		}
	}

	ans := 0
	for _, a := range cnt {
		if len(a) == 0 {
			continue
		}

		// 存在空的情况
		a = append(a, []int{0, 0}...)
		sort.Ints(a)
		length := len(a)

		// 只看最长串情况(包含-2)
		// a[length-1]-1是二者l0 l1相同情况
		// a[length-2] 可能有第三个值为0的情况
		// 只看最短串情况(其他也满足)
		ans = max(ans, max(max(a[length-1]-2, min(a[length-1]-1, a[length-2])), a[length-3]))
	}

	if ans == 0 {
		return -1
	}
	return ans
}
