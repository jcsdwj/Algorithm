package leetCode

import "fmt"

// 436. 寻找右区间

// 给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。

// 区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。注意 i 可能等于 j 。

// 返回一个由每个区间 i 的 右侧区间 在 intervals 中对应下标组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。

func TestLeetCode436() {
	intervals1 := [][]int{{1, 2}}
	intervals2 := [][]int{{3, 4}, {2, 3}, {1, 2}}
	intervals3 := [][]int{{1, 4}, {2, 3}, {3, 4}}
	intervals4 := [][]int{{4, 4}}
	fmt.Println(findRightInterval(intervals1))
	fmt.Println(findRightInterval(intervals2))
	fmt.Println(findRightInterval(intervals3))
	fmt.Println(findRightInterval(intervals4))
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)

	// 存在且最小
	ans := []int{}
	for i := 0; i < n; i++ {
		var temp = -1
		endi := intervals[i][1]
		for j := 0; j < n; j++ {
			startj := intervals[j][0]
			if startj >= endi {
				if temp == -1 {
					temp = j
					continue
				}
				startTemp := intervals[temp][0]
				if startj < startTemp {
					temp = j
				}
			}
		}
		ans = append(ans, temp)
	}
	return ans
}
