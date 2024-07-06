/**
  @author: 伍萬
  @since: 2024/6/26
  @desc: //TODO
**/

package leetCode

import "fmt"

// 2741. 特别的排列
// n个不同的数

func TestLeetCode2741() {
	nums1 := []int{2, 3, 6}
	nums2 := []int{1, 4, 3}
	fmt.Println(specialPerm2(nums1))
	fmt.Println(specialPerm2(nums2))
}

// nums数组会变化(代码哪里写错了)
func specialPerm(nums []int) int {
	sum := 0
	// 选中一个数放第1位
	var cur = 0
	for i := 0; i < len(nums); i++ {
		cur = nums[i]
		sum += countSpe(append(nums[:i], nums[i+1:]...), cur)
	}
	return sum
}

func countSpe(nums []int, cur int) int {
	if len(nums) == 0 {
		return 1
	}

	sum := 0
	// 从nums中找满足条件的
	var temp = 0
	for i := 0; i < len(nums); i++ {
		temp = nums[i]
		if nums[i]%cur != 0 || cur%nums[i] != 0 {
			sum += countSpe(append(nums[:i], nums[i+1:]...), temp)
		}
	}
	return sum
}

const mod = 1e9 + 7

// 记忆化搜索
func specialPerm2(nums []int) int {
	n := len(nums)
	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(state, i int) int {
		if f[state][i] != -1 {
			return f[state][i]
		}
		if state == (1 << i) {
			return 1
		}
		f[state][i] = 0
		for j := 0; j < n; j++ {
			if i == j || state>>j&1 == 0 {
				continue
			}
			if nums[i]%nums[j] != 0 && nums[j]%nums[i] != 0 {
				continue
			}
			f[state][i] = (f[state][i] + dfs(state^(1<<i), j)) % mod
		}
		return f[state][i]
	}

	res := 0
	for i := 0; i < n; i++ {
		res = (res + dfs((1<<n)-1, i)) % mod
	}
	return res
}

// 状态压缩动态规划
func specialPerm3(nums []int) int {
	n := len(nums)
	f := make([][]int64, 1<<n)
	for i := range f {
		f[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		f[1<<i][i] = 1
	}

	for state := 1; state < (1 << n); state++ {
		for i := 0; i < n; i++ {
			if state>>i&1 == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if i == j || state>>j&1 == 0 {
					continue
				}
				x := nums[i]
				y := nums[j]
				if x%y != 0 && y%x != 0 {
					continue
				}
				f[state][i] = (f[state][i] + f[state^(1<<i)][j]) % mod
			}
		}
	}

	var sum int64
	for i := 0; i < n; i++ {
		sum = (sum + f[(1<<n)-1][i]) % mod
	}
	return int(sum)
}
