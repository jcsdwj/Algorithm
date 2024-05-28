package leetCode

import (
	"fmt"
	"sort"
)

// 315. 计算右侧小于当前元素的个数

func TestLeetCode315() {
	nums1 := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(nums1))
}

// 超时
func countSmaller(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		count := 0
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		dp[i] = count
	}
	return dp
}

// 依然超时
func countSmaller2(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	dict := make(map[int]int)
	dict[nums[n-1]]++
	for i := n - 2; i >= 0; i-- {
		for k, v := range dict {
			if nums[i] > k {
				dp[i] += v
			}
		}
		dict[nums[i]]++
	}
	return dp
}

var dp []int

// 二分查找
func countSmaller3(nums []int) []int {
	n := len(nums)

	res := []int{}
	for i := n - 1; i >= 0; i-- {
		x := binary(nums[i])
		res = append([]int{x}, res...)
		// dp=append()
	}
	return res
}

func binary(target int) int {
	left := 0
	right := len(dp)
	for left < right {
		mid := left + (right-left)/2
		if dp[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 树状数组写法
var a, c []int

func countSmaller4(nums []int) []int {
	resultList := []int{}
	discretization(nums)
	c = make([]int, len(nums)+5)
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(nums[i])
		resultList = append(resultList, query(id-1))
		update(id)
	}
	for i := 0; i < len(resultList)/2; i++ {
		resultList[i], resultList[len(resultList)-1-i] = resultList[len(resultList)-1-i], resultList[i]
	}
	return resultList
}
func lowBit(x int) int {
	return x & (-x)
}

func update(pos int) {
	for pos < len(c) {
		c[pos]++
		pos += lowBit(pos)
	}
}

func query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += c[pos]
		pos -= lowBit(pos)
	}
	return ret
}

func discretization(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	a = make([]int, 0, len(nums))
	for num := range set {
		a = append(a, num)
	}
	sort.Ints(a)
}

func getId(x int) int {
	return sort.SearchInts(a, x) + 1
}

// 归并排序写法
