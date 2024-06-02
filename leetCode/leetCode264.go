/**
  @author: 伍萬
  @since: 2024/6/2
  @desc: //TODO
**/
//264. 丑数 II

package leetCode

import "fmt"

func TestLeetCode264() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(1))
}

//264. 丑数 II

// 暴力枚举（超时）
func nthUglyNumber(n int) int {
	var j = 1
	for i := 0; i < n; i++ {
		for !isUgly(j) {
			j++
		}
		j++
	}
	j--
	return j
}

// 判断是否为丑数
func isUgly(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return isUgly(n / 2)
	}
	if n%3 == 0 {
		return isUgly(n / 3)
	}
	if n%5 == 0 {
		return isUgly(n / 5)
	}
	return false
}

// 由2 3 5 构成的数
// 构建 2 3 5 丑数列表
// 下个丑数 从更小的丑数乘出来的
func nthUglyNumber2(n int) int {
	// 作为头指针（都指向当前最小丑数，如果等于最小丑数了，就乘以2或3或5）
	p2 := 1
	p3 := 1
	p5 := 1

	// 作为头结点的值（当前最小丑数乘以2 3 5）
	product2, product3, product5 := 1, 1, 1

	// 丑数集合
	ugly := make([]int, n+1)
	// 丑数指针(个数)
	p := 1

	for p <= n {
		// 获取最小值
		minValue := min264(product2, product3, product5)
		ugly[p] = minValue
		p++ // 向后移动1位

		if minValue == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if minValue == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}
		if minValue == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}
	}

	return ugly[n]
}

// 获取三个数里最小的
func min264(a, b, c int) int {
	temp := min(a, b)
	return min(temp, c)
}
