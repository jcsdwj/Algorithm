package leetCode

import "Algorithm/common"

// 143. 重排链表

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 1 2 3 4 -> 1 4 2 3
// 1 2 3 4 -> 1 5 2 4 3

func reorderList(head *common.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := middleNode(head)
	left := head
	right := mid.Next
	mid.Next = nil
	right = reverseList(right)
	mergeList(left, right)
}

// 翻转链表
func reverseList(right *common.ListNode) *common.ListNode {
	prev := right
	if prev.Next == nil {
		// 只要一个节点情况
		return prev
	}

	cur := prev.Next
	for cur.Next != nil {
		// 找到最后一个节点
		cur = cur.Next
		prev = prev.Next
	}
	prev.Next = nil
	cur.Next = reverseList(right)
	return cur
}

// 找到中间节点 快慢指针
func middleNode(head *common.ListNode) *common.ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 合并链表
func mergeList(left, right *common.ListNode) {
	var leftTmp, rightTmp *common.ListNode
	for left != nil && right != nil {
		leftTmp = left.Next
		rightTmp = right.Next
		left.Next = right
		left = leftTmp

		right.Next = left
		right = rightTmp
	}
}
