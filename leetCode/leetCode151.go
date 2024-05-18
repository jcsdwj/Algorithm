package leetCode

import (
	"fmt"
	"strings"
)

// 2024.05.15

/*
151. 反转字符串中的单词
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
*/

func TestLeetCode151() {
	s1 := "a good   example"
	fmt.Println(reverseWords(s1))
}

func reverseWords(s string) string {
	str := strings.Split(strings.TrimSpace(s), " ")
	ansStr := []string{}
	for i := len(str) - 1; i >= 0; i-- {
		if strings.Contains(str[i], " ") || str[i] == "" {
			continue
		}
		ansStr = append(ansStr, strings.TrimSpace(str[i]))
	}
	return strings.Join(ansStr, " ")
}