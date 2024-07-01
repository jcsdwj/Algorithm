/**
  @author: 伍萬
  @since: 2024/6/22
  @desc: //TODO
**/

package leetCode

import "fmt"

func TestLeetCode2663() {
	fmt.Println(smallestBeautifulString("abcz", 26))
	fmt.Println(smallestBeautifulString("dc", 4))
}

// 2663. 字典序最小的美丽字符串
// 由字母表前k个字母组成，不包含任何长度为2或更长的回文子串(不能连续相同)
// 字典序大于s
// 当前已经是美丽字符串
// 任何长度为 2 或者更长的回文字符串，都包含一个长度为 2 或者 3 的回文字符串 本身就是美丽字符串
// 不能和前一个和前两个相同
func smallestBeautifulString(s string, k int) string {
	for i := len(s) - 1; i >= 0; i-- {
		blockedCharacters := make(map[byte]bool)
		for j := 1; j < 3; j++ {
			// i的前两位确定不变
			if i-j >= 0 {
				blockedCharacters[s[i-j]] = true
			}
		}
		for j := 1; j < 4; j++ {
			// 在k范围内且不等于i的前两位 因为字典序要大 所以从三个值选一个肯定有满足条件的 例如 前两个是ab 第三个为c就好
			// 找到一个满足的就达到要求
			if int(s[i]-'a')+j+1 <= k && !blockedCharacters[s[i]+byte(j)] {
				return generate(s, i, j)
			}
		}
	}
	return ""
}

// index为要修改的下标 offset为字符增大偏移量
func generate(s string, index int, offset int) string {
	res := []byte(s)
	res[index] += byte(offset)

	// 后面直接abc开始
	// 修改后的字符不能和前两个相同
	for i := index + 1; i < len(s); i++ {
		blockedCharacters := make(map[byte]bool)
		for j := 1; j < 3; j++ {
			if i-j >= 0 {
				blockedCharacters[res[i-j]] = true
			}
		}

		// 为什么是abc
		for c := byte('a'); c <= byte('c'); c++ {
			if !blockedCharacters[c] {
				res[i] = c
				break
			}
		}
	}
	return string(res)
}

func smallestBeautifulString2(s string, k int) string {
	limit := 'a' + byte(k) // 限制位 不能到这个点
	strs := []byte(s)
	n := len(strs)
	i := n - 1 // 从最后一个字母开始
	strs[i]++  // 先加一

	for i < n {
		if strs[i] == limit { // 需要进位
			if i == 0 { // 无法进位
				return ""
			}
			// 进位
			strs[i] = 'a'
			i--
			strs[i]++ // 前面的加一位
		} else if i > 0 && strs[i] == strs[i-1] || i > 1 && strs[i] == strs[i-2] {
			// 等于前两位 形成回文串
			strs[i]++ // 如果 s[i] 和左侧的字符形成回文串，就继续增加 s[i]
		} else {
			// 因为自会和前面的比 所以当前的改好需要让后面的值也参与判断 没问题会一直++
			i++ // 反过来检查后面是否有回文串
		}
	}

	return string(strs)
}
