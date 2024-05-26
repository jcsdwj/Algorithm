/**
  @author: 伍萬
  @since: 2024/5/25
  @desc: //TODO
**/

package leetCode

import "sort"

// 493. 翻转对
// i < j 且 nums[i] > 2*nums[j]
// 超时
func reversePairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if nums[i] > 2*nums[j] {
				count++
			}
		}
	}
	return count
}

// 递归也超时
func reversePairs2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[0] > 2*nums[i] {
			count++
		}
	}
	return count + reversePairs2(nums[1:])
}

// 归并排序
func reversePairs3(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// []int(nil)和[]int{}的区别 前者是nil的切片形式 后者已经分配内存了
	n1 := append([]int(nil), nums[:n/2]...)
	n2 := append([]int(nil), nums[n/2:]...)

	cnt := reversePairs(n1) + reversePairs(n2) // 递归完毕后，n1 和 n2 均为有序

	// 统计重要翻转对 (i,j) 的数量
	// 由于 n1 和 n2 均为有序，可以用两个指针同时遍历
	j := 0
	for _, v := range n1 { // n1中索引肯定小于n2
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		cnt += j
	}

	// n1 和 n2 归并填入 nums
	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt

}

// 树状数组
func reversePairs4(nums []int) (cnt int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	allNums := make([]int, 2*n)
	for _, v := range nums {
		allNums = append(allNums, v, 2*v) // 所有值翻倍
	}

	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i < 2*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	t := newFenwickTree(k)
	for i, v := range nums {
		// 统计之前插入了多少个比2*v大的数
		cnt += i - t.sum(kth[2*v])
		t.add(kth[v])
	}
	return
}

type fenwick struct {
	tree []int
}

func newFenwickTree(n int) fenwick {
	return fenwick{make([]int, n+1)}
}
func (f fenwick) add(i int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i]++
	}
}

func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}
