package leetCode

import (
	"Algorithm/common"
	"sort"
)

// 230. 二叉搜索树中第K小的元素

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）

var nums230 []int

// 二叉搜索树是中序遍历
func kthSmallest(root *common.TreeNode, k int) int {
	nums230 = []int{}
	traverse(root)
	sort.Ints(nums230)
	return nums230[k-1]
}

func traverse(root *common.TreeNode) {
	if root == nil {
		return
	}
	nums230 = append(nums230, root.Val)
	traverse(root.Left)
	traverse(root.Right)
}
