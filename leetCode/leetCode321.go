package leetCode

// 321. 拼接最大数

// 给你两个整数数组 nums1 和 nums2，它们的长度分别为 m 和 n。数组 nums1 和 nums2 分别代表两个数各位上的数字。同时你也会得到一个整数 k。

// 请你利用这两个数组中的数字中创建一个长度为 k <= m + n 的最大数，在这个必须保留来自同一数组的数字的相对顺序。

// 返回代表答案的长度为 k 的数组。

// 保证当前值最大

// 单调栈
// 未理解
func maxSubsequence(a []int, k int) (s []int) {
	for i, v := range a {
		// k-len(s)<= len(a)-i-1
		for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
			// 剩余元素的数量多于s中可以存放的(后面还有) 且当前值大于栈顶元素 删除栈顶元素
			s = s[:len(s)-1]
		}

		if len(s) < k {
			s = append(s, v)
		}
	}
	return
}

// a数组是否比b数组小?
func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b) // b后面还有值
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			// b减少一个元素
			merged[i], b = b[0], b[1:]
		} else {
			// a减少一个元素，因为a大
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	// nums1至少选start个

	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubsequence(nums1, i)   // 从nums1选i个
		s2 := maxSubsequence(nums2, k-i) // 从nums2选k-i个
		merged := merge(s1, s2)

		// 如果新组成的数组大于当前的
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}
