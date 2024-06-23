/**
  @author: 伍萬
  @since: 2024/6/23
  @desc: //TODO
**/

package leetCode

// 520. 检测大写字母

func detectCapitalUse(word string) bool {
	// 所有字母都是大写
	// 都不是大写
	// 只有首字母大写
	countBig := 0
	countSmall := 0
	firstIsBig := false
	for i := 0; i < len(word); i++ {
		if word[i] <= 'Z' && word[i] >= 'A' {
			if i == 0 {
				firstIsBig = true
			}
			countBig++
		}
		if word[i] <= 'z' && word[i] >= 'a' {
			countSmall++
		}
	}
	if countBig == len(word) || countSmall == len(word) {
		return true
	}
	if countBig == 1 && firstIsBig {
		return true
	}
	return false
}
