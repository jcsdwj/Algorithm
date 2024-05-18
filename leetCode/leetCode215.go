package leetCode

import "sort"

// 215. 数组中的第K个最大元素

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 直接排序求解
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 快速选择算法
func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

// 得到k位置的值(从小到大)
func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}

	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j++ {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		// j作为尾部
		return quickselect(nums, l, j, k)
	} else {
		// j+1作为头部
		return quickselect(nums, j+1, r, k)
	}
}

// 堆排序
func findKthLargest3(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// 构建堆
func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
