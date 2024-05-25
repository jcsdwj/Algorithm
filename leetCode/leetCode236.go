package leetCode

import "Algorithm/common"

// 236. 二叉树的最近公共祖先

// 递归方法
// root为根节点的情况
// 1. 左右子树分别包含p q
// 2. p或q为root节点（并且包含另一个节点）
func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	// 没找到
	if root == nil {
		return nil
	}

	// 如果p或q为root则返回root
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)   // 左子树找最近公共节点
	right := lowestCommonAncestor(root.Right, p, q) // 右子树找最近公共节点

	// 左右两边各找到一个
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

// 深度搜索法

func lowestCommonAncestor2(root, p, q *common.TreeNode) *common.TreeNode {
	parent := map[int]*common.TreeNode{}
	visited := map[int]bool{}

	var dfs func(*common.TreeNode)
	dfs = func(r *common.TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r // 记录父节点
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true // 访问过该节点
		p = parent[p.Val]     // 更新为父节点
	}
	for q != nil {
		if visited[q.Val] { // 若访问过该节点 说明是相同父节点
			return q
		}
		q = parent[q.Val]
	}

	return nil

}
