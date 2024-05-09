/**
  @author: 伍萬
  @since: 2024/5/10
  @desc: //TODO
**/

package nowcoder

import "sort"

// 颜色分类

func sortColor(colors []int) []int {
	// write code here
	sort.Ints(colors)
	return colors
}

// 方法2
// 直接统计个数
func sortColor2(colors []int) []int {
	// write code here
	red := 0
	white := 0
	for i := 0; i < len(colors); i++ {
		if colors[i] == 0 {
			red++
		} else if colors[i] == 1 {
			white++
		}
	}

	for i := 0; i < len(colors); i++ {
		if i < red {
			colors[i] = 0
		} else if i >= red && i < red+white {
			colors[i] = 1
		} else {
			colors[i] = 2
		}
	}
	return colors
}
