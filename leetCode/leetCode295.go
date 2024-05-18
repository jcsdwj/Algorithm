package leetCode

import "sort"

// 295. 数据流的中位数

// 涉及堆，有待继续研究

// 中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。

// 例如 arr = [2,3,4] 的中位数是 3 。
// 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
// 实现 MedianFinder 类:

// MedianFinder() 初始化 MedianFinder 对象。

// void addNum(int num) 将数据流中的整数 num 添加到数据结构中。

// double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。

// 超出输出限制
type MedianFinder struct {
	nums   []int
	length int
}

func (this *MedianFinder) FindMedian() float64 {
	if this.length%2 == 0 {
		return float64(this.nums[this.length/2])
	}
	return (float64(this.nums[this.length/2]) + float64(this.nums[this.length/2-1])) / 2
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)
	this.length++
	sort.Ints(this.nums)
}

func Constructor() MedianFinder {
	return MedianFinder{
		nums:   []int{},
		length: 0,
	}
}

// 字典写法(超时)
type MedianFinder2 struct {
	dict   map[int]int
	length int   // 总共有多少个数
	keys   []int // key集合
}

func Constructor2() MedianFinder2 {
	return MedianFinder2{
		dict:   make(map[int]int),
		length: 0,
		keys:   make([]int, 0),
	}
}

func (this *MedianFinder2) AddNum2(num int) {
	this.dict[num]++
	this.length++
	if this.dict[num] == 1 {
		this.keys = append(this.keys, num)
	}
	sort.Ints(this.keys)
}

func (this *MedianFinder2) FindMedian2() float64 {
	if this.length == 1 || len(this.keys) == 1 {
		return float64(this.keys[0])
	}

	temp := this.length / 2 // 中位数位置
	sum := 0
	var v1, v2 int
	// 奇数找到temp和temp-1的值 1,2,3
	for i := 0; i < len(this.keys); i++ {
		sum += this.dict[this.keys[i]]
		if sum > temp {
			// 找到v2的值
			v2 = this.keys[i]
			// 判断v1是和v2相同还是前一个数
			if sum-this.dict[this.keys[i]] == temp {
				v1 = this.keys[i-1]
			} else {
				v1 = this.keys[i]
			}
			break
		}
	}

	if this.length%2 == 0 {
		return float64(float64(v1)+float64(v2)) / 2
	}
	return float64(v2)
}

// 优先队列解法

type MedianFinder3 struct {
	queMin, queMax hp
}

// sort.IntSlice用来建堆
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
