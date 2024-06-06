/**
  @author: 伍萬
  @since: 2024/6/6
  @desc: //TODO
**/

package leetCode

import "fmt"

// 2938. 区分黑球与白球
// 1移到右侧 0移到左侧的最小步数
// 每次只能移动1步

func TestLeetCode2938() {
	//fmt.Println(minimumSteps("101"))  // 1
	//fmt.Println(minimumSteps("100"))  // 2
	fmt.Println(minimumSteps("0111")) // 0
}

func minimumSteps(s string) int64 {
	// 统计0 和 1 出现次数，判断最右是否都为1
	countZero := 0
	countOne := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			countOne++
		} else {
			countZero++
		}
	}

	// [0:countZero] 全为0
	// [countZero:len(s)]全为1

	startZero := 0
	startOne := len(s) - 1
	countOp := 0
	for startZero < countZero && startOne >= countZero {
		for s[startZero] == '0' {
			startZero++
		} // 从左往右找到第一个1
		for s[startOne] == '1' {
			startOne--
		} // 从右往左找到第一个0
		if startZero < startOne {
			countOp += startOne - startZero
			// 移动元素 startZero和startOne
			s = s[:startZero] + "0" + s[startZero+1:startOne] + "1" + s[startOne+1:]
		}
	}
	return int64(countOp)
}

// 贪心法
// 遍历字符串的时候每碰到一个 0 就贪心的将其往左交换直到它最终的位置
// 遍历到这个 0 时，因为它左边的 0 已经都交换到最终位置了
// 所以它的左边是一串连续的 1，那么只要加上遍历时碰到 1 的个数即可
// 例如110110 第1个0移到1左边就是 011110 第2个0 会遇到4个1
func minimumSteps2(s string) int64 {
	ans := int64(0)
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			sum++
		} else {
			// 遇到0把它移到1右边的位置去
			ans += int64(sum)
		}
	}
	return ans
}
