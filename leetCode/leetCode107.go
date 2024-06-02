/**
  @author: 伍萬
  @since: 2024/6/2
  @desc: //TODO
**/

package leetCode

// 107. 二叉树的层序遍历 II

import (
	"Algorithm/common"
)

// 遍历一层后存储该层节点
func levelOrderBottom(root *common.TreeNode) [][]int {
	levelOrder := [][]int{}
	if root == nil {
		return levelOrder
	}

	queue := []*common.TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		level := []int{}
		size := len(queue) // 当前层有多少节点
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level) // 放入一层的节点
	}

	// 从下往上遍历
	for i := 0; i < len(levelOrder)/2; i++ {
		levelOrder[i], levelOrder[len(levelOrder)-1-i] = levelOrder[len(levelOrder)-1-i], levelOrder[i]
	}

	return levelOrder
}
