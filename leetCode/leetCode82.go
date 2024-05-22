package leetCode

import (
	"Algorithm/common"
	"sort"
)

// 82. 删除排序链表中的重复元素 II

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表

func deleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	l := head
	dict := make(map[int]int)
	nums := []int{}
	for l != nil {
		dict[l.Val]++
		l = l.Next
	}

	for k, v := range dict {
		if v == 1 {
			nums = append(nums, k)
		}
	}
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)

	// 构建链表
	newHead := &common.ListNode{
		Val:  nums[0],
		Next: nil,
	}

	l = newHead
	for i := 1; i < len(nums); i++ {
		temp := &common.ListNode{
			Val:  nums[i],
			Next: nil,
		}
		l.Next = temp
		l = l.Next
	}

	return newHead
}
