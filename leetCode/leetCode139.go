package leetCode

import (
	"fmt"
	"strings"
)

// 139. 单词拆分

func TestLeetCode139() {
	wordDict1 := []string{"leet", "code"}
	wordDict2 := []string{"apple", "pen"}
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak2("leetcode", wordDict1))
	fmt.Println(wordBreak2("applepenapple", wordDict2))
	fmt.Println(wordBreak2("catsandog", wordDict3))
}

// 递归法超时
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}

	var flag = false

	for i := 0; i < len(wordDict); i++ {
		str := wordDict[i]
		if strings.HasPrefix(s, wordDict[i]) {
			length := len(str)
			flag = wordBreak(s[length:], wordDict)
		}
		if flag {
			return flag
		}
	}
	return flag
}

func wordBreak2(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1) // 表示前i个字符串是否能被wordDict构成
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}
	for i := 1; i <= n; i++ {
		if dict[s[:i]] {
			dp[i] = true
			continue
		}

		// 遍历访问过的
		for j := i - 1; j >= 0; j-- {
			// 前j个元素可以且s[j:i]在word里
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
