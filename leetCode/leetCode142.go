/**
  @author: 伍萬
  @since: 2024/5/26
  @desc: //TODO
**/

package leetCode

import (
	"Algorithm/common"
	"math"
)

//124. 二叉树中的最大路径和

// 这种写法肯定会包含当前节点 有问题
func maxPathSum(root *common.TreeNode) int {

	if root == nil {
		return 0
	}
	val := root.Val
	rightVal := maxPathSum(root.Right) // 左右子树最大路径和
	leftVal := maxPathSum(root.Left)
	if val >= 0 {
		if rightVal > 0 {
			val += rightVal
		}
		if leftVal > 0 {
			val += rightVal
		}
	}
	if val < 0 {
		// 多种情况
		val = max(val, max(leftVal, max(rightVal, max(leftVal+val, max(rightVal+val, leftVal+val+rightVal)))))
	}
	return val
}

// [0] 会返回42 不知道为什么
var max142 = math.MinInt

func maxPathSum2(root *common.TreeNode) int {
	maxGain(root)
	return max142
}

// 包含当前节点最大值
func maxGain(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	leftVal := maxGain(root.Left)
	rightVal := maxGain(root.Right)
	curVal := root.Val

	// 只能选择一边进行返回 或者当前节点直接返回
	curMax := max(curVal+leftVal, max(curVal+rightVal, curVal))

	// 中间存在经过左根右的情况 但不返回 因为不可能再经过上面的节点了
	max142 = max(max(curMax, max142), curVal+rightVal+leftVal)

	return curMax
}

// 第二种写法没问题
func maxPathSum3(root *common.TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*common.TreeNode) int
	maxGain = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
