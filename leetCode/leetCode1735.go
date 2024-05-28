package leetCode

// 1735. 生成乘积数组的方案数

// 长度为0 乘积为2
func waysToFillArray(queries [][]int) []int {
	return nil
}

// 长度与乘积的关系
func computeRelation(a, b int) int {
	if a == 1 {
		return 1
	}
	if a == 2 {
		count := 0
		for i := 1; i <= b; i++ {
			if b%i == 0 {
				count++
			}
		}
		return count
	}
	return 0
}

// 1 1种
// 假设k有x个因子 要构成n个参数
