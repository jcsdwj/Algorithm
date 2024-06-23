/**
  @author: 伍萬
  @since: 2024/6/17
  @desc: //TODO
**/

package leetCode

//522. 最长特殊序列 II

func findLUSlength522(strs []string) int {
	// 判断是否为子字符串
	ans := -1
	for i, s := range strs {
		x := len(s)
		for j, t := range strs {
			// 是子字符串
			if i != j && isSubSeq(s, t) {
				x = -1
				break
			}
		}
		ans = max(ans, x)
	}
	return ans
}

// 判断a是否为b的子字符串
func isSubSeq(a, b string) bool {
	if a == b {
		return true
	}

	m, n := len(a), len(b)
	i := 0
	for j := 0; i < m && j < n; j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == m
}
