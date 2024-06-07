package leetCode

import (
	"fmt"
	"sort"
)

// 3072. 将元素分配到两个数组中 II

func LeetCode3072() {
	nums1 := []int{2, 1, 3, 3}
	nums2 := []int{5, 14, 3, 1, 2}
	nums3 := []int{3, 3, 3, 3}
	fmt.Println(resultArray(nums1))
	fmt.Println(resultArray(nums2))
	fmt.Println(resultArray(nums3))
}

// 传统方法超时
func resultArray(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}

	// greaterCount(arr,val) 严格大于val的元素个数
	for i := 2; i < len(nums); i++ {
		value1 := greaterCount(arr1, nums[i])
		value2 := greaterCount(arr2, nums[i])
		if value1 > value2 {
			arr1 = append(arr1, nums[i])
		} else if value1 < value2 {
			arr2 = append(arr2, nums[i])
		} else {
			l1 := len(arr1)
			l2 := len(arr2)
			if l1 < l2 {
				arr1 = append(arr1, nums[i])
			} else if l1 > l2 {
				arr2 = append(arr2, nums[i])
			} else {
				arr1 = append(arr1, nums[i])
			}
		}
	}

	return append(arr1, arr2...)
}

func greaterCount(arr []int, val int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > val {
			count++
		}
	}

	return count
}

// 方法2：使用树状数组方法
func resultArray2(nums []int) []int {
	n := len(nums)
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	index := make(map[int]int)
	for i, num := range sortedNums {
		index[num] = i + 1
	}

	arr1, arr2 := []int{nums[0]}, []int{nums[1]}
	tree1, tree2 := NewBinaryIndexedTree(n), NewBinaryIndexedTree(n)
	tree1.Add(index[nums[0]])
	tree2.Add(index[nums[1]])

	for i := 2; i < n; i++ {
		count1 := len(arr1) - tree1.Get(index[nums[i]])
		count2 := len(arr2) - tree2.Get(index[nums[i]])
		if count1 > count2 || (count1 == count2 && len(arr1) <= len(arr2)) {
			arr1 = append(arr1, nums[i])
			tree1.Add(index[nums[i]])
		} else {
			arr2 = append(arr2, nums[i])
			tree2.Add(index[nums[i]])
		}
	}

	return append(arr1, arr2...)
}

type BinaryIndexedTree struct {
	tree []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
	return &BinaryIndexedTree{tree: make([]int, n+1)}
}

func (bit *BinaryIndexedTree) Add(i int) {
	for i < len(bit.tree) {
		bit.tree[i]++
		i += i & -i
	}
}

func (bit *BinaryIndexedTree) Get(i int) int {
	sum := 0
	for i > 0 {
		sum += bit.tree[i]
		i -= i & -i
	}
	return sum
}
