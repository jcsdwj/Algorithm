package leetCode

// 133. 克隆图

import (
	"Algorithm/common"
)

// 访问过的节点
var visitedNode map[*common.Node]*common.Node

func cloneGraph(node *common.Node) *common.Node {
	visitedNode = make(map[*common.Node]*common.Node)
	return dfs133(node)
}

// DFS方法
func dfs133(node *common.Node) *common.Node {
	if node == nil {
		return node
	}

	// 如果被访问过了
	if _, ok := visitedNode[node]; ok {
		return visitedNode[node]
	}

	// 克隆节点
	cloneNode := &common.Node{
		Val:       node.Val,
		Neighbors: nil,
	}
	visitedNode[node] = cloneNode

	for _, n := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, dfs133(n))
	}
	return cloneNode
}

// BFS方法
func bfs133(node *common.Node) *common.Node {
	if node == nil {
		return node
	}

	queue := []*common.Node{node}
	// 保存新构建的节点
	visitedNode[node] = &common.Node{
		Val:       node.Val,
		Neighbors: []*common.Node{},
	}
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		for _, neighbors := range n.Neighbors {
			if _, ok := visitedNode[neighbors]; ok {
				// 放入哈希表和队列
				visitedNode[neighbors] = &common.Node{neighbors.Val, []*common.Node{}}
				queue = append(queue, neighbors)
			}
			// 更新当前节点的邻居表
			visitedNode[n].Neighbors = append(visitedNode[n].Neighbors, visitedNode[neighbors])
		}
	}
	return visitedNode[node]
}
