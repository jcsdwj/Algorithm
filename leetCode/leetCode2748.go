/**
  @author: 伍萬
  @since: 2024/6/20
  @desc: //TODO
**/

package leetCode

import "fmt"

// 2748. 美丽下标对的数目

func TestLeetCode2748() {
	fmt.Println(countBeautifulPairs([]int{2, 5, 1, 4}))
	fmt.Println(countBeautifulPairs([]int{11, 21, 12}))
}

// nums[i]的第一个数和nums[j]的最后一个数互质
func countBeautifulPairs(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] >= 10 {
			nums[i] /= 10
		}
		for j := i + 1; j < n; j++ {
			if gcd(nums[i], nums[j]%10) == 1 {
				res++
			}
		}
	}
	return res
}

// 判断两个数是否互质
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
