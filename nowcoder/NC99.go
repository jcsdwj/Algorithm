/**
  @author: 伍萬
  @since: 2024/3/3
  @desc: //TODO
**/

package nowcoder

import (
	"Algorithm/common"
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 树的直径
 * @param n int整型 树的节点个数
 * @param Tree_edge Interval类一维数组 树的边
 * @param Edge_value int整型一维数组 边的权值
 * @return int整型
 */

func solve(n int, Tree_edge []*common.Interval, Edge_value []int) int {
	// write code here
	path := [][]int{}
	//used := make([]bool, n)
	for i := 0; i < n; i++ {
		path = append(path, make([]int, n))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				path[i][j] = 0
			} else {
				path[i][j] = math.MaxInt
			}
		}
	}

	// 赋值
	for i := 0; i < len(Edge_value); i++ {
		start := Tree_edge[i].Start
		end := Tree_edge[i].End
		path[start][end] = Edge_value[i]
		path[end][start] = Edge_value[i]
	}
	return 0
}

//先从任意一点出发找到离它最远的一个直径点x,
//再从找到的最远直径点x出发找到离当前点最远的直径点y,
//这样x到y就是这个树上的直径

type Node struct {
	start int
	edges map[int]int
}

// 转换函数
func trans(Tree_edge []*common.Interval, Edge_value []int) map[int]*Node {
	length := len(Tree_edge)
	for i := 0; i < length; i++ {

	}
	return nil
}
