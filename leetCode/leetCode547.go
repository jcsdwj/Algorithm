package leetCode

import "fmt"

// 547. 省份数量

func TestLeetCode547() {
	isConnected1 := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected1))
	isConnected2 := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected2))
}

// 深度遍历所有相关的节点
// 访问过的设为0
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected) // n个城市
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				count++
				dfs547(isConnected, i)
			}
		}
	}

	return count
}

func dfs547(isConnected [][]int, start int) {
	// 起始节点置空
	isConnected[start][start] = 0
	n := len(isConnected)
	// 遍历与start相连的点
	for i := 0; i < n; i++ {
		if isConnected[start][i] == 1 {
			isConnected[start][i] = 0
			isConnected[i][start] = 0 // 这条边访问过了
			dfs547(isConnected, i)
		}
	}
}

// 官方深度搜索
func findCircleNum2(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true // 访问过的设为true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			// 没有访问过
			ans++
			dfs(i)
		}
	}
	return
}

// 官方广度搜索
func findCircleNum3(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	for i, v := range vis {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				vis[from] = true // 访问过设为true
				for to, conn := range isConnected[from] {
					if conn == 1 && !vis[to] {
						// 有连接且没访问过
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return
}
