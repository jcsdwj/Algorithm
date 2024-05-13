package leetCode

/*
189. 轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/
// 类似轮转链表

func rotate(nums []int, k int) {
	k = k % len(nums)
	temp := append(nums[len(nums)-k:len(nums)], nums[:len(nums)-k]...)
	copy(nums, temp)
}
