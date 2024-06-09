/**
  @author: 伍萬
  @since: 2024/6/9
  @desc: //TODO
**/

package nowcoder

import "Algorithm/common"

// ONT9 牛牛的递增之旅
func removeDuplicates(head *common.ListNode) *common.ListNode {
	// write code here
	slow := head

	if slow == nil {
		return nil
	}
	fast := head.Next
	if fast == nil {
		return head
	}

	for fast != nil {
		// 快指针等于慢指针
		if fast.Val == slow.Val {
			// 快指针向后移
			fast = fast.Next
		} else {
			// 遇到不同的 慢指针的Next直接指向fast
			// 快慢指针都向后移
			slow.Next = fast
			slow = slow.Next
			fast = fast.Next
		}
	}
	slow.Next = fast
	return head
}
