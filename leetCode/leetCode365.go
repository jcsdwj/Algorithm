package leetCode

import "fmt"

// 365. 水壶问题

// 有两个水壶，容量分别为 x 和 y 升。水的供应是无限的。确定是否有可能使用这两个壶准确得到 target 升。

// 你可以：
// 装满任意一个水壶
// 清空任意一个水壶
// 将水从一个水壶倒入另一个水壶，直到接水壶已满，或倒水壶已空。

func TestLeetCode365() {
	fmt.Println(canMeasureWater(3, 5, 4))
	fmt.Println(canMeasureWater(2, 6, 5))
	fmt.Println(canMeasureWater(1, 2, 3))
}

func canMeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}

	var max, min int
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}

	dictYu := make(map[int]bool)
	dictYu[0] = true
	for i := 1; i < max; i++ {
		dictYu[(i*min)%max] = true
	}
	return dictYu[target%max]
}
