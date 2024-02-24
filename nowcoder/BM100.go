/**
  @author: 伍萬
  @since: 2024/2/16
  @desc: //TODO
**/

package nowcoder

import "container/list"

type Solution struct {
	// write code here
	cap   int
	lst   *list.List
	cache map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) Solution {
	// write code here
	return Solution{
		capacity,
		list.New(),
		map[int]*list.Element{},
	}
}

func (this *Solution) get(key int) int {
	// write code here
	e := this.cache[key]
	if e == nil {
		return -1
	}
	this.lst.MoveToFront(e) // 移动到最前面
	return e.Value.(entry).value
}

func (this *Solution) set(key int, value int) {
	// write code here
	e := this.cache[key]
	if e != nil {
		// 更新
		e.Value = entry{key, value}
		this.lst.MoveToFront(e)
		return
	}

	// 插入到最前
	this.cache[key] = this.lst.PushFront(entry{key, value})
	if len(this.cache) > this.cap {
		// delete删除键值对
		delete(this.cache, this.lst.Remove(this.lst.Back()).(entry).key)
	}
}
