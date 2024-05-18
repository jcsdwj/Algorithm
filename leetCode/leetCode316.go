package leetCode

import (
	"fmt"
	"strings"
)

// 316. 去除重复字母

// 给你一个字符串s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的
// 字典序最小（要求不能打乱其他字符的相对位置）。

// 2024-05-16

func TestLeetCode316() {
	fmt.Println(removeDuplicateLetters("bcabc")) // abc
}

// 字典序最小
func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			// 如果没用出现过
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}

// 单调栈
func removeDuplicateLetters2(s string) string {
	stk := []rune{}

	// 维护一个计数器记录字符串中字符的数量
	// 因为输入为 ASCII 字符，大小 256 够用了
	count := [256]int{}
	for _, c := range s {
		count[c]++
	}

	inStack := [256]bool{} // 字符是否在栈中
	for _, c := range s {
		// 每遍历过一个字符，都将对应的计数减一
		count[c]--

		if inStack[c] {
			// 重复字符忽略
			continue
		}

		for len(stk) > 0 && stk[len(stk)-1] > c {
			// 若之后不存在栈顶元素了，则停止 pop
			if count[stk[len(stk)-1]] == 0 {
				break
			}
			// 若之后还有，则可以 pop(放入当前更小的)
			inStack[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, c)
		inStack[c] = true
	}

	sb := strings.Builder{}
	for i := 0; i < len(stk); i++ {
		sb.WriteRune(stk[i])
	}
	return sb.String()
}
