/**
  @author: 伍萬
  @since: 2024/6/13
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"sort"
)

// 2813. 子序列最大优雅度

// [1,2] [2,1] [3,1]

// 第一个值是利润
// 优雅度total_profit + distinct_categories2
// total_profit子序列所有项目利润总和
// distinct_categories所选子序列所含类别中不同类别的数量
// 长度为k的子序列，从里面挑出k个
// 先选利润最大的 选好后从不同的里面选利润最大的
// 如果都是相同的 就挑选相同里利润最大的
// 排序

func TestLeetCode2813() {
	items1 := [][]int{
		{3, 2}, {5, 1}, {10, 1},
	}
	items2 := [][]int{
		{3, 1}, {3, 1}, {2, 2}, {5, 3},
	}
	items3 := [][]int{
		{1, 1}, {2, 1}, {3, 1},
	}
	items4 := [][]int{{1, 1}, {4, 1}}
	items5 := [][]int{{2, 2}, {8, 3}, {8, 3}}
	fmt.Println(findMaximumElegance(items1, 2))
	fmt.Println(findMaximumElegance(items2, 3))
	fmt.Println(findMaximumElegance(items3, 3))
	fmt.Println(findMaximumElegance(items4, 1))
	fmt.Println(findMaximumElegance(items5, 2))
}

// 要考虑取相同还是不同
func findMaximumElegance(items [][]int, k int) int64 {
	dict := make(map[int]bool)
	temp := [][]int{}

	var maxIndex int
	var maxValue int

	// 先选个最大利润的
	for i := 0; i < len(items); i++ {
		v1 := items[i][0]
		if v1 > maxValue {
			maxValue = v1
			maxIndex = i
		}
	}
	dict[items[maxIndex][1]] = true // 存入最大值
	temp = append(temp, items[maxIndex])
	items = append(items[:maxIndex], items[maxIndex+1:]...)

	for len(temp) != k {
		maxValue = items[0][0] // 假设第0个值是最大利润
		maxIndex = 0
		for i := 1; i < len(items); i++ {
			// 找最大不同值
			v1 := items[i][0]        // 第i个值的利润
			v2 := items[i][1]        // 第i个值的类型
			v3 := items[maxIndex][1] // 当前最大值的类型

			// 两个都在dict中 比较最大值
			if dict[v2] && dict[v3] {
				if v1 > maxValue {
					maxValue = v1
					maxIndex = i
				}
			} else if dict[v2] && !dict[v3] {
				// 若i在maxIndex不在
				continue
			} else if !dict[v2] && dict[v3] {
				maxValue = v1
				maxIndex = i
			} else if !dict[v2] && !dict[v3] {
				// 两个都不在 比利润
				if v1 > maxValue {
					maxValue = v1
					maxIndex = i
				}
			}
		}
		temp = append(temp, items[maxIndex])
		dict[items[maxIndex][1]] = true
		items = append(items[:maxIndex], items[maxIndex+1:]...)
	}

	var sum int64
	for i := 0; i < len(temp); i++ {
		sum += int64(temp[i][0])
	}

	return sum + int64(len(dict))*int64(len(dict))
}

/*
按照profit排序，选前k个可以保证利润最大 但不保证种类最大
如果第k+1个的种类不同 则需要判断替换掉哪种
*/
func findMaximumElegance2(items [][]int, k int) int64 {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	}) // 排序 递减情况

	categoryMap := make(map[int]bool)
	var res, profit int64
	var st []int
	for i, item := range items {
		if i < k {
			profit += int64(item[0])
			// 出现过
			if categoryMap[item[1]] {
				st = append(st, item[0])
			} else {
				categoryMap[item[1]] = true
			}
		} else if !categoryMap[item[1]] && len(st) > 0 {
			// 没有出现过并且前面有重复的 从最后一个重复的替换
			profit += int64(item[0] - st[len(st)-1])
			st = st[:len(st)-1]
			categoryMap[item[1]] = true
		}
		res = maxInt64(res, profit+int64(len(categoryMap)*len(categoryMap)))
	}
	return res
}
