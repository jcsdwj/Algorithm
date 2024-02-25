/**
  @author: 伍萬
  @since: 2024/2/24
  @desc: //TODO
**/

package nowcoder

import (
	"Algorithm/common"
	"math"
)

func judgeIt(root *common.TreeNode) []bool {
	// write code here
	flag1 := isSearchTree(root, math.MinInt, math.MaxInt)
	flag2 := isAllTree(root)
	return []bool{flag1, flag2}
}

// 判断是否为搜索树
func isSearchTree(root *common.TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if root.Val <= left || root.Val >= right {
		return false
	}

	return isSearchTree(root.Left, left, root.Val) && isSearchTree(root.Right, root.Val, right)
}

// 这是层序遍历 得到的数字排序不满足递增 因为左子树所有的值都比根节点的小
func isAllTree(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	var flag = true
	stack := []*common.TreeNode{}
	stack = append(stack, root)

	var i = 0

	// 遍历树
	for i < len(stack) {
		temp := stack[i]

		if temp == nil {
			i++
			continue
		}

		if temp.Left != nil {
			stack = append(stack, temp.Left)

		} else {
			stack = append(stack, nil)
		}

		if temp.Right != nil {
			stack = append(stack, temp.Right)
		} else {
			stack = append(stack, nil)
		}
		i++
	}

	for j := 0; j < len(stack)-1; j++ {
		if stack[j] == nil && stack[j+1] != nil {
			flag = false
			break
		}
	}

	return flag
}
