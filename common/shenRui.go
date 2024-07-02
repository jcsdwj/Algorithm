/**
  @author: 伍萬
  @since: 2024/6/26
  @desc: //TODO
**/

package common

// 判断链表是否有环
func isNotCircle(head *ListNode) bool {
	// 用字典存储遍历过的节点
	dict := make(map[*ListNode]bool)
	l := head
	for l != nil {
		// 表示已经访问过该节点
		if dict[l] == true {
			return false
		}
		dict[l] = true
		l = l.Next
	}
	return true
}
