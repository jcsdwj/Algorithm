/**
  @author: 伍萬
  @since: 2024/6/1
  @desc: //TODO
**/

package leetCode

import "fmt"

// 2928. 给小朋友们分糖果 I

// 每个小朋友的糖果不超过limit
// n个糖果分给3个小朋友

func TestLeetCode2928() {
	fmt.Println(distributeCandies(5, 2))
	fmt.Println(distributeCandies(3, 3))
}

// 枚举法
func distributeCandies(n int, limit int) int {
	count := 0
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if n-i-j <= limit && n-i-j >= 0 {
				fmt.Println(i, j, n-i-j)
				count++
			}
		}
	}
	return count
}

// 优化枚举法
// 第一个小朋友有x个 剩余n-x
// 如果n-x>2*limit 不满足条件
// 如果n-x<=2*limit 则一个小朋友得分
// 根据不等式，一个小朋友至少得分n-x-limit个糖果(如果大于0) 才能保证另一个不超过limit
// 同时他最多能拿n-x颗(同时得不大于limit)
func distributeCandies2(n int, limit int) int {
	ans := 0
	for i := 0; i <= min(limit, n); i++ {
		if n-i > 2*limit {
			continue
		}

		// 这两个值范围中
		ans += min(n-i, limit) - max(0, n-i-limit) + 1
	}
	return ans
}

// 容斥原理
func cal(x int) int {
	if x < 0 {
		return 0
	}
	return x * (x - 1) / 2
}

func distributeCandies3(n int, limit int) int {
	return cal(n+2) - 3*cal(n-limit+1) + 3*cal(n-(limit+1)*2+2) - cal(n-3*(limit+1)+2)
}
