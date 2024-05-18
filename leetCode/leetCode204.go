/**
  @author: 伍萬
  @since: 2024/5/18
  @desc: //TODO
**/
//204. 计数质数

package leetCode

import "fmt"

func TestLeetCode204() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(0))
	fmt.Println(countPrimes(1))
}

// 超时
func countPrimes(n int) int {
	isPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime[i] = false
				continue
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

// 利用乘法做
func countPrimes2(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i < n/2+1; i++ {
		for j := 2; j*i < n; j++ {
			isPrime[i*j] = false
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
