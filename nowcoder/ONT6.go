/**
  @author: 伍萬
  @since: 2024/5/9
  @desc: //TODO
**/
/**
旋转链表
*/
package nowcoder

import "Algorithm/common"

func rotateLinkedList(head *common.ListNode, k int) *common.ListNode {
	// write code here
	if head == nil {
		return head
	}
	count := 0 // 节点数量
	l := head
	for l != nil {
		count++
		l = l.Next
	}

	if k%count == 0 {
		return head
	}
	k = k % count
	after := head
	for i := 1; i < count-k; i++ {
		after = after.Next // 得到右移末尾节点的位置
	}
	before := after.Next // 作为新头结点的位置
	after.Next = nil
	tail := before
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = head
	return before
}
