package leetCode

import "fmt"

// 797. 所有可能的路径

// 0到n的所有路径

func TestLeetCode797() {
	// graph1 := [][]int{
	// 	{1, 2}, {3}, {3}, {},
	// }
	//graph2 := [][]int{
	//	{4, 3, 1}, {3, 2, 4}, {3}, {4}, {},
	//}
	//fmt.Println(allPathsSourceTarget(graph1))
	// fmt.Println(allPathsSourceTarget(graph2))

	graph3 := [][]int{
		{3, 1}, {4, 6, 7, 2, 5},
		{4, 6, 3}, {6, 4}, {7, 6, 5},
		{6}, {7}, {},
	}
	fmt.Println(allPathsSourceTarget(graph3))
}

var res797 [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	res797 = [][]int{}
	dfs797(graph, 0, len(graph)-1, []int{})
	return res797
}

// 判断cur能否到末尾节点
func dfs797(graph [][]int, cur int, last int, path []int) {
	path = append(path, cur) // 添加当前节点
	if cur == last {
		// 这里得用深拷贝 数组的值变了temp里的也会变
		temp := make([]int, len(path))
		copy(temp, path)
		// 当前节点是最后一个
		res797 = append(res797, temp)
		return
	}

	curPath := graph[cur] // 当前节点到哪些节点有路径
	for i := 0; i < len(curPath); i++ {
		dfs797(graph, curPath[i], last, path)
	}
}
