/**
  @author: 伍萬
  @since: 2024/2/25
  @desc: //TODO
**/

package nowcoder

import (
	"Algorithm/common"
	"math"
)

func maxPathSum(root *common.TreeNode) int {
	// write code here
	res = math.MinInt
	getMax(root)
	return res
}

var res int

// 以root节点为起点的最大值
func getMax(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	left := int(math.Max(float64(getMax(root.Left)), float64(0)))
	right := int(math.Max(float64(getMax(root.Right)), float64(0)))

	priceNewpath := root.Val + left + right
	res = int(math.Max(float64(priceNewpath), float64(res)))

	return root.Val + int(math.Max(float64(left), float64(right)))
}
