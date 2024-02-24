/**
  @author: 伍萬
  @since: 2024/2/14
  @desc: //TODO
**/

package nowcoder

import (
	"fmt"
	"math"
	"strings"
)

// 滑动窗口
func TestBM90() {
	var s1 = "wyqaalfdtavrmkvrgbrmauenibfxrzjpzzqzlveexayjkvdsyueboloymtvfugvtgnutkyzhaztlvwdvqkwgvmejhbpdimwqbslnrkutkpehnkefwblrprcxvtaffzxitivmssgehklvwqastojihmhcfkhnlexemtrhnxlujxrgpuyiikspycufodubisfwnydaxrwhqqpfkppuzjlzlfhbjbcttkriixkiohpexgjjvafxjqyvyfyjhbccltlvsvdgeumdavoyxtvhmtekzctidxkqsxmlvrrzmefobtmznhizdmlcoemudwkvuyirscqegvsjrfkgoshrgsvvyhrbgdycehtwjlcrjucabpgsjnjqhhnfqeiwhgalptjyflpoiuqjjwdslpiswvxobfljnnwdhgdortezpulysoqddbxbwuqabdjqqhtzpxpjsvkjrvhjmzoralvzhlzkqkbgrwijvzspvcymafymfmfhaaghnfsdrvmlruuntmcqqbdqideprkxrmfbanvfeqrphnlwjxbzqcegmhnczxbslitnvotaemroadkjclksppzeyoiznlsytnopchritiyvlleqypiqgjugxeikpclipzxtgoebxcxkpdaoulecuajueretvpbkqbgwrkaooxbeaduvoaxlaifgblzwdcjtfpsxbsnrrovturokrovtycbcqcytfjomygj"
	var s2 = "baxtr"
	fmt.Println(minWindow2(s1, s2))
}

// 获取最小字符串（时间复杂度过高）
func minWindow(S string, T string) string {

	dictS := make(map[string]int)
	dictT := make(map[string]int)

	// 获取字典
	for i := 0; i < len(S); i++ {
		dictS[S[i:i+1]]++
	}
	for j := 0; j < len(T); j++ {
		dictT[T[j:j+1]]++
	}

	if !checkWindow(dictS, dictT) {
		return ""
	}

	return getMinWindow(S, T, dictS, dictT)
}

// 得到最小字符串
func getMinWindow(S string, T string, dictS map[string]int, dictT map[string]int) string {
	if len(S) == len(T) {
		return S
	}

	firstStr := S[:1]
	lastStr := S[len(S)-1:]
	var lstr = S
	var rstr = S

	// 获取删除第一个字符的最小字符串
	dictS[firstStr]--
	if checkWindow(dictS, dictT) {
		lstr = getMinWindow(S[1:], T, dictS, dictT)
	}

	if checkWindow(dictS, dictT) {
		rstr = getMinWindow(S[:len(S)-1], T, dictS, dictT)
	}
	dictS[lastStr]++

	if len(lstr) > len(rstr) {
		return rstr
	}
	return lstr
}

// 判断是否能完全包含
func checkWindow(dictS map[string]int, dictT map[string]int) bool {
	for k, v := range dictT {
		if dictS[k] < v {
			return false
		}
	}
	return true
}

func minWindow2(S string, T string) string {
	if strings.Contains(S, T) {
		return T
	}

	dictS := make(map[string]int)
	dictT := make(map[string]int)

	// 统计出现频率
	//for i := 0; i < len(S); i++ {
	//	dictS[S[i:i+1]]++
	//}
	for j := 0; j < len(T); j++ {
		dictT[T[j:j+1]]++
	}

	var left = 0
	var right = 0

	// 起始位置
	var start = 0
	// 最小长度
	var minLen = math.MaxInt

	// 判断是否完全匹配
	var match int

	for right < len(S) {
		c1 := S[right : right+1]
		if dictT[c1] != 0 {
			// 需要这个字符
			dictS[c1]++
			if dictS[c1] == dictT[c1] {
				match++
			}
		}
		right++

		for match == len(dictT) {
			// 作为候选
			if right-left < minLen {
				minLen = right - left
				start = left
			}

			// 找到要使用的左指针位置  忽略掉不使用的
			c2 := S[left : left+1]
			if dictT[c2] != 0 {
				dictS[c2]--
				if dictS[c2] < dictT[c2] {
					match--
				}
			}
			left++
		}
	}

	if minLen == math.MaxInt {
		return ""
	}
	return S[start : start+minLen]
}
