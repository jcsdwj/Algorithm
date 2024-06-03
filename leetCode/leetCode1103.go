package leetCode

import (
	"fmt"
	"math"
)

// 1103. 分糖果 II

func TestLeetCode1103() {
	// fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
}

// 直接模拟
func distributeCandies1103(candies int, num_people int) []int {
	nums := make([]int, num_people)
	start := 1
	for candies > 0 {
		// 正常要放入的糖果数
		if candies < start {
			nums[(start-1)%num_people] += candies
			candies = 0
		} else {
			nums[(start-1)%num_people] += start
			candies -= start
			start++
		}
	}
	return nums
}

// 通过等差数列求和公式可以知道第几个值会超过糖果数(x)
func distributeCandies1103_2(candies int, num_people int) []int {
	n := num_people
	// how many people received complete gifts
	p := int(math.Sqrt(float64(2*candies)+0.25) - 0.5)
	remaining := candies - int(float64((p+1)*p)*0.5)
	rows, cols := p/n, p%n

	d := make([]int, n)
	for i := 0; i < n; i++ {
		// complete rows
		d[i] = (i+1)*rows + int(float64(rows*(rows-1)*n)*0.5)
		// cols in the last row
		if i < cols {
			d[i] += i + 1 + rows*n
		}
	}
	// remaining candies
	d[cols] += remaining
	return d
}
