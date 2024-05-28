package leetCode

import "sort"

// 406. 根据身高重建队列
// 有多少个身高大于等于自己
func reconstructQueue(people [][]int) [][]int {
	// 最矮的人为i people[i][1]为a 则它前面得有a个人
	// 第二矮的人为j people[j][1]为b b不可能大于a b肯定小于等于a
	// 如果等于a 则放在i后面 否则放在i前面 位置为b+1

	// 先从低到高考虑 比a高的比比b高的人更多
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	ans := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1 //矮的人的位置
		for i := range ans {
			if ans[i] == nil {
				// 该位置为空
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}

// 从高到低考虑
func reconstructQueue2(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return
}
