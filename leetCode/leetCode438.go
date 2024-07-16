/**
  @author: 伍萬
  @since: 2024/7/9
  @desc: //TODO
**/

package leetCode

// 438. 找到字符串中所有字母异位词

// 从某个索引开始后几位是异位词
func findAnagrams(s string, p string) []int {
	length := len(p)
	start, end := 0, 0

	dictP := make(map[string]int)
	dictS := make(map[string]int)

	// 统计p的词频
	for i := 0; i < length; i++ {
		dictP[p[i:i+1]]++
	}

	valid := 0
	res := []int{}
	for end < len(s) {
		c := s[end : end+1]
		end++

		// 进行窗口内数据的一系列更新
		if _, ok := dictP[c]; ok {
			// 包含这个词
			dictS[c]++
			if dictP[c] == dictS[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for end-start >= length {
			// 当窗口符合条件时，把起始索引加入 res
			if valid == len(dictP) {
				res = append(res, start)
			}
			d := s[start : start+1]
			start++
			//窗口数据更新
			if _, ok := dictP[d]; ok {
				if dictS[d] == dictP[d] {
					valid--
				}
				dictS[d]--
			}
		}
	}
	return res
}
