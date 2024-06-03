package leetCode

import "fmt"

// 207. 课程表

// 先修课程[a,b] 表示b到a有路径 想到a必须经过b
// 存不存在[a,b],[a,c]，有两个先修课的情况
// 主要是不能形成环
// 思路没问题(但还是有问题,理解问题)

func TestLeetCode207() {
	// pre1 := [][]int{
	// 	{1, 0},
	// }
	pre2 := [][]int{
		{1, 0}, {0, 1},
	}
	// fmt.Println(canFinish(2, pre1))
	fmt.Println(canFinish(2, pre2))

	pre3 := [][]int{
		{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14},
		{13, 1}, {15, 1}, {17, 4},
	}
	fmt.Println(canFinish(20, pre3))
}

var visited map[int]bool

func canFinish(numCourses int, prerequisites [][]int) bool {
	dict := make(map[int]bool)
	for i := 0; i < len(prerequisites); i++ {
		dict[prerequisites[i][0]] = true
		dict[prerequisites[i][1]] = true
	}

	if len(dict) < numCourses {
		return false
	}

	edges := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		edges[i] = make([]int, numCourses)
	}
	for i := 0; i < len(prerequisites); i++ {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		edges[b][a] = 1
	}
	visited = make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		// 第i个节点开始查
		ans := dfs207(edges, i)
		if !ans {
			return false
		}
	}
	return true
}

// 深度搜索判断有没有环
func dfs207(edges [][]int, cur int) bool {
	// 访问过该节点返回false
	if visited[cur] {
		return false
	}

	visited[cur] = true
	for i := 0; i < len(edges); i++ {
		if i == cur {
			continue
		}
		if edges[cur][i] == 1 {
			ans := dfs207(edges, i)
			if !ans {
				return false
			}
		}
	}

	// 回退
	visited[cur] = false

	return true
}
