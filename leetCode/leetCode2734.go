/**
  @author: 伍萬
  @since: 2024/6/27
  @desc: //TODO
**/

package leetCode

import "fmt"

//2734. 执行子串操作后的字典序最小字符串

func TestLeetCode2734() {
	fmt.Println(smallestString3("cbabc"))
	fmt.Println(smallestString3("acbbc"))
	fmt.Println(smallestString3("leetcode"))
	fmt.Println(smallestString3("ba"))  // aa
	fmt.Println(smallestString3("aab")) // 应该是aaa
	fmt.Println(smallestString3("aa"))  // az
}

// 前一个字符替换当前 连续的子字符串
func smallestString(s string) string {
	// 找到第一个a的位置
	// a在0 修改后面的
	// a在中间 修改前面的
	// a在最后 修改前面的
	// a不存在 全部修改
	str := ""
	index := -1 // a第一次出现的位置
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			if index == -1 {
				index = i
			}
			count++
		}
	}

	if count == len(s) {
		return s[:len(s)-1] + "z"
	}

	if index == -1 {
		// a不存在
		for i := 0; i < len(s); i++ {
			str += string(s[i] - 1)
		}
	} else if index == 0 {
		str += "a" // 第一位先加上a
		for i := 1; i < len(s); i++ {
			if s[i] == 'a' {
				str += s[i:]
				break // 直接退出
			} else {
				str += string(s[i] - 1)
			}
		}
	} else if index == len(s)-1 {
		for i := 0; i < len(s); i++ {
			if s[i] == 'a' {
				str += "a"
			} else {
				str += string(s[i] - 1)
			}
		}
	} else {
		// 中间情况
		for i := 0; i < len(s); i++ {
			if s[i] == 'a' {
				str += s[i:]
				break // 直接退出
			} else {
				str += string(s[i] - 1)
			}
		}
	}

	return str
}

// 连续的变换操作 （超时）
func smallestString2(s string) string {
	if len(s) == 0 {
		return s
	}
	// a为第一位
	// a为最后一位
	// a在中间
	// 连续的a
	if s[0] == 'a' && len(s) > 1 {
		// 首位为a的情况
		return "a" + smallestString(s[1:])
	}
	return changeStr(s)
}

// 转换
func changeStr(s string) string {
	if len(s) == 1 && s[0] == 'a' {
		// a作为最后一个字符
		return "z"
	}
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			str += s[i:] // 后面的直接加
			break
		}
		str += string(s[i] - 1)
	}
	return str
}

// 找到第一个非a字符修改它  因为只有a->z会导致字典序变大 其他都是变小
// 超时(可能是操作切片的问题)
func smallestString3(s string) string {
	index1 := findIndexFirstNotA(s)
	if index1 == len(s) {
		// 说明全是a 最后一个改成z
		return s[:len(s)-1] + "z"
	}

	index2 := findIndexFirstNotAAfterA(s, index1)
	str := s[:index1]
	for i := index1; i < index2; i++ {
		str += string(s[i] - 1)
	}
	return str + s[index2:]
}

func findIndexFirstNotA(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != 'a' {
			return i
		}
	}
	return len(s)
}

// 找到非A后的第一个A（停止更新）
func findIndexFirstNotAAfterA(s string, index int) int {
	for i := index + 1; i < len(s); i++ {
		if s[i] == 'a' {
			return i
		}
	}
	return len(s)
}
