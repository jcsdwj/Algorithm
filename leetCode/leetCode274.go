package leetCode

import (
	"fmt"
	"sort"
)

/*
274. H 指数
给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，
并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。
*/

func TestLeetCode274() {
	citations1 := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations1))
}

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	if len(citations) == 1 {
		if citations[0] >= 1 {
			return 1
		} else {
			return 0
		}
	}
	sort.Ints(citations)
	dict := make(map[int]int) // 统计出现频次

	for i := len(citations) - 1; i >= 0; i-- {
		dict[citations[i]]++

		// 比当前小也加1
		for j := 0; j < i; j++ {
			dict[j]++
		}
	}

	for i := len(citations) - 1; i >= 0; i-- {
		if dict[citations[i]] >= citations[i] {
			return citations[i]
		}
	}
	return len(citations)
}
