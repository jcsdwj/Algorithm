/**
  @author: 伍萬
  @since: 2024/2/16
  @desc: //TODO
**/

package nowcoder

import (
	"container/list"
	"fmt"
	"math"
)

// 1 表示set
// 2 表示get 返回此结果
// k 表示存放记录数
// 移除调用次数最少 有相同值则移除最近未使用

func TestBM101() {
	nums := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 3},
		{1, 4, 4},
		{2, 4},
		{2, 3},
		{2, 2},
		{2, 1},
		{1, 5, 5},
		{2, 4},
	}

	//nums := [][]int{
	//	{1, 1, 1},
	//	{1, 2, 2},
	//	{1, 3, 2},
	//	{1, 2, 4},
	//	{1, 3, 5},
	//	{2, 2},
	//	{1, 4, 4},
	//	{2, 1},
	//}

	fmt.Println(LFU(nums, 4))
}

type Cache struct {
	capacity  int                   // 容量
	lst       *list.List            // data值
	dict      map[int]*list.Element // key映射值
	countDict map[int]int           // 统计调用频率
	minFreq   int                   // 最小频率
}

type data struct {
	key   int
	value int
}

func (c *Cache) Get(key int) int {
	// 更新minFreq
	var t = math.MaxInt
	for _, v := range c.countDict {
		if v < t {
			t = v
		}
	}
	c.minFreq = t

	e := c.dict[key]
	if e == nil {
		return -1
	}

	// 使用频率加1
	c.countDict[key]++

	if c.countDict[key] < c.minFreq {
		c.minFreq = c.countDict[key]
	}

	c.lst.MoveToFront(e) // 也移动到前面
	return e.Value.(data).value
}

func (c *Cache) Set(key int, value int) {
	// 更新minFreq
	var t = math.MaxInt
	for _, v := range c.countDict {
		if v < t {
			t = v
		}
	}
	c.minFreq = t

	e := c.dict[key]

	// 更新直接返回
	if e != nil {
		e.Value = data{
			key:   key,
			value: value,
		}
		c.lst.MoveToFront(e)
		c.countDict[key]++

		// 更新最小使用频率
		if c.countDict[key] < c.minFreq {
			c.minFreq = c.countDict[key]
		}
		return
	}

	// 插入
	c.dict[key] = c.lst.PushFront(data{key, value})
	c.countDict[key]++ // 统计频率 新放入的不统计频率

	if len(c.dict) > c.capacity {
		// 从后遍历lst 找到第一个频率等于minFreq的删除
		for v := c.lst.Back(); v != nil; v = v.Prev() {
			k := v.Value.(data).key
			if k == key {
				// 插入当前key 跳过
				continue
			}

			if c.countDict[k] == c.minFreq {
				delete(c.dict, c.lst.Remove(v).(data).key)
				delete(c.countDict, k)
				return
			}
		}
	}
}

func LFU(operators [][]int, k int) []int {

	cache := Cache{
		capacity:  k,
		lst:       list.New(),
		dict:      map[int]*list.Element{},
		countDict: make(map[int]int),
		minFreq:   math.MaxInt,
	}

	// write code here
	answer := []int{}
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			cache.Set(operators[i][1], operators[i][2])
		} else {
			value := cache.Get(operators[i][1])
			answer = append(answer, value)
		}
	}

	return answer
}
