package leetCode

import "fmt"

// 3067. 在带权树网络中统计可连接服务器对数目

func TestLeetCode3067() {
	edges1 := [][]int{
		{0, 1, 1}, {1, 2, 5}, {2, 3, 13},
		{3, 4, 9}, {4, 5, 2},
	}
	edges2 := [][]int{
		{0, 6, 3}, {6, 5, 3}, {0, 3, 1},
		{3, 2, 7}, {3, 1, 6}, {3, 4, 2},
	}
	fmt.Println(countPairsOfConnectableServers(edges1, 1))
	fmt.Println(countPairsOfConnectableServers(edges2, 3))
}

// 服务器i可连接的服务器对数
// 获取所有节点之间的长度
// 需要构成图 利用乘法原理(但怎么保证求余为0呢)
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges)
	ans := make([]int, n)

	// 更新节点路径长度
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := i + 1; k < j; j++ {
				if edges[i][j] == 0 && edges[i][k] != 0 && edges[k][j] != 0 {
					edges[i][j], edges[j][i] = edges[i][k]+edges[k][j], edges[i][k]+edges[k][j]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				// 必须有路径
				if edges[j][k] == 0 {
					continue
				}

				// 不能等于c节点
				if j == i || k == i {
					continue
				}

				if edges[i][j]%signalSpeed != 0 {
					continue
				}
				if edges[i][k]%signalSpeed != 0 {
					continue
				}

				// a b 没有公共边
				if edges[j][i]+edges[j][k] == edges[i][k] || edges[k][i]+edges[j][k] == edges[i][j] {
					continue
				}
				count++
			}
		}
		ans[i] = count
	}
	return ans
}

// 根枚举
func countPairsOfConnectableServers2(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	graph := make([][][]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], []int{v, w})
		graph[v] = append(graph[v], []int{u, w})
	}

	var dfs func(int, int, int) int
	dfs = func(p, root, curr int) int {
		res := 0
		if curr == 0 {
			res++
		}
		for _, e := range graph[p] {
			v, cost := e[0], e[1]
			if v != root {
				res += dfs(v, p, (curr+cost)%signalSpeed)
			}
		}
		return res
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		pre := 0
		for _, e := range graph[i] {
			v, cost := e[0], e[1]
			cnt := dfs(v, i, cost%signalSpeed)
			res[i] += pre * cnt
			pre += cnt
		}
	}
	return res
}
