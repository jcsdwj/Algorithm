package leetCode

import (
	"container/heap"
	"fmt"
	"sort"
)

func TestLeetCode373() {
	nums1_1 := []int{1, 7, 11}
	nums2_1 := []int{2, 4, 6}
	fmt.Println(kSmallestPairs(nums1_1, nums2_1, 3))
	nums1_2 := []int{1, 1, 2}
	nums2_2 := []int{1, 2, 3}
	fmt.Println(kSmallestPairs(nums1_2, nums2_2, 2))
}

// 373. 查找和最小的 K 对数字

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res Res
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			temp := []int{nums1[i], nums2[j]}
			res = append(res, temp)
		}
	}

	sort.Sort(res)
	res = res[:k]
	return res
}

// 超内存
type Res [][]int

func (x Res) Len() int { return len(x) }
func (x Res) Less(i, j int) bool {
	if x[i][0]+x[i][1] < x[j][0]+x[j][1] {
		return true
	}
	return false
	// if x[i][0] < x[j][0] {
	// 	return true
	// }
	// if x[i][0] == x[j][0] && x[i][1] <= x[j][1] {
	// 	return true
	// }
	// return false
}
func (x Res) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// 采用堆的方法
func kSmallestPairs2(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	ans := make([][]int, 0, min(k, n*m)) // 预分配空间
	h := hp2{{nums1[0] + nums2[0], 0, 0}}
	for len(h) > 0 && len(ans) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j == 0 && i+1 < n {
			heap.Push(&h, tuple{nums1[i+1] + nums2[0], i + 1, 0})
		}
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

// 涉及heap的使用
type tuple struct{ s, i, j int }
type hp2 []tuple

// 实现堆的五个接口
func (h hp2) Len() int           { return len(h) }
func (h hp2) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp2) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
