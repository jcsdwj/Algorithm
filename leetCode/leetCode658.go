package leetCode

import "sort"

// 658. 找到 K 个最接近的元素

// 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。

// 整数 a 比整数 b 更接近 x 需要满足：

// |a - x| < |b - x| 或者
// |a - x| == |b - x| 且 a < b

func findClosestElements(arr []int, k int, x int) []int {
	// 稳定排序
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
