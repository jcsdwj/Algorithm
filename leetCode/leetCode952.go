package leetCode

// 952. 按公因数计算最大组件大小

// 给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：

// 有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
// 只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
// 返回 图中最大连通组件的大小 。

// 连通图最多节点数

// 并查集方法 貌似还有BFS方法
type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return unionFind{parent, make([]int, n)}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) merge(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}

func largestComponentSize(nums []int) (ans int) {
	m := 0
	for _, num := range nums {
		m = max(m, num)
	}
	uf := newUnionFind(m + 1)
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				uf.merge(num, i)
				uf.merge(num, num/i)
			}
		}
	}
	cnt := make([]int, m+1)
	for _, num := range nums {
		rt := uf.find(num)
		cnt[rt]++
		ans = max(ans, cnt[rt])
	}
	return
}
