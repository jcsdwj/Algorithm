package leetCode

import (
	"fmt"
	"slices"
	"strings"
)

// 936. 戳印序列

// 数组由每个回合中被印下的最左边字母的索引组成
// 未理解

func TestLeetCode936() {
	stamp1 := "abc"
	target1 := "ababc"
	fmt.Println(movesToStamp(stamp1, target1))
	stamp2 := "abca"
	target2 := "aabcaca"
	fmt.Println(movesToStamp(stamp2, target2))

	stamp3 := "oz"
	target3 := "ooozz"
	fmt.Println(movesToStamp(stamp3, target3))
}

func movesToStamp(stamp string, target string) []int {

	// 一定要包含stamp
	if !strings.Contains(target, stamp) {
		return nil
	}

	if stamp == target {
		return []int{0}
	}

	// 如果包含印章 最后一次盖
	// 如果印章包含这个戳前面的 且包含戳后面的 可以盖
	// 假如包含的位置是k 则 stamp[0:len(stamp)-1] 要包含 0:k stamp[1:] 要包含k+m:

	nums := []int{}

	m := len(stamp)
	n := len(target)

	for i := 0; i <= n-m; i++ {
		//印章
		if target[i:i+m] == stamp {
			nums = append(nums, i)
			i += m - 1 // 移动最后一个字符的位置
		}
	}
	// 循环后保存最后几次需要盖章的地方

	// 进行校验 首中尾
	var temp = []int{}

	// 印章位置切分
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		if i == 0 {
			// 计算头部
			if start == 0 {
				// 最后印章在0
				continue
			}

			// 看要求包含即可
			if strings.HasPrefix(stamp, target[:start]) {
				// 必须以这个开头
				temp = append(temp, 0)
			} else {
				return nil
			}
		}
		if i == len(nums)-1 {
			// 计算尾部
			if start == n-m {
				// 最后印章在末尾
				continue
			}
			if strings.HasSuffix(stamp, target[start+m:]) {
				// 在末尾以这个字符串结尾
				temp = append(temp, n-m)
			} else {
				return nil
			}
		}

		// 非头非尾
		if i != 0 && i != len(nums)-1 {
			// 中间情况包含即可
			end := nums[i+1]
			if strings.Contains(stamp, target[start:end]) {
				// 子字符串在印章中的位置 abcd abc bcd
				index := strings.Index(stamp, target[start:end])
				temp = append(temp, start+m-index)

			} else {
				return nil
			}
		}
	}

	return append(temp, nums...)
}

// 逆向操作
func movesToStamp2(stamp string, target string) []int {
	n := len(stamp)
	s := []byte(target)
	var res []int

	// start为序列的起始位置
	var check = func(start int) {
		for i := 0; i < len(stamp); i++ {
			j := i + start

			// 盖不了
			if stamp[i] != s[j] && s[j] != '#' {
				return
			}
		}
		flag := false
		for i := 0; i < len(stamp); i++ {
			j := i + start
			if s[j] != '#' {
				flag = true
			}
			s[j] = '#'
		}
		if flag {
			// 添加印章
			res = append(res, start)
		}
	}

	// 为什么要正反各校验一次(存在先盖后再盖前的操作?)

	// 正着盖
	for i := 0; i+n-1 < len(s); i++ {
		check(i)
	}
	// 反着盖
	for i := len(s) - n; i >= 0; i-- {
		// 后面的印章盖不到
		check(i)
	}
	for i := range s {
		if s[i] != '#' {
			return []int{}
		}

	}
	slices.Reverse(res)
	return res

}
