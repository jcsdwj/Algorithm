/**
  @author: 伍萬
  @since: 2024/6/16
  @desc: //TODO
**/

package leetCode

import "fmt"

//521. 最长特殊序列 Ⅰ

func TestLeetCode521() {
	fmt.Println(findLUSlength("aba", "cdc"))
	fmt.Println(findLUSlength("aaa", "bbb"))
	fmt.Println(findLUSlength("aaa", "aaa"))
}

func findLUSlength(a string, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	// 没有出现过的集合
	dict1 := make(map[string]int)
	dict2 := make(map[string]int)
	for i := 0; i < len(a); i++ {
		dict1[a[i:i+1]]++
	}
	for i := 0; i < len(b); i++ {
		dict2[b[i:i+1]]++
	}

	var maxLen1 = 0
	var maxLen2 = 0
	for k, v := range dict1 {
		if dict2[k] == 0 {
			maxLen1 += v
		}
	}
	for k, v := range dict2 {
		if dict1[k] == 0 {
			maxLen2 += v
		}
	}
	v := max(maxLen1, maxLen2)
	if v == 0 {
		return -1
	}
	return v
}
