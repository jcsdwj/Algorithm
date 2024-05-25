package leetCode

import "Algorithm/common"

//103. 二叉树的锯齿形层序遍历

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
//（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）

func zigzagLevelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 用栈保存节点
	// 从栈顶元素开始遍历
	// 获取层数 隔一层数组翻转
	res := [][]int{}
	stack := []*common.TreeNode{root}
	levelSize := 1
	temp := []int{}
	level := 0

	for len(stack) != 0 {
		// 当前层的节点全部访问
		for i := 0; i < levelSize; i++ {
			p := stack[i]
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			temp = append(temp, p.Val)
		}
		stack = stack[levelSize:] // 出队列
		levelSize = len(stack)    // 获取当前层有多少节点数

		if level%2 == 1 {
			// 逆序
			reverseArray(temp)
		}

		level++ // 更新层数

		res = append(res, temp)
		temp = []int{} //清空数据
	}

	return res
}

// 翻转数组
func reverseArray(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
