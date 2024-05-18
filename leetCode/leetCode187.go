package leetCode

import "fmt"

// 187. 重复的DNA序列

// DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.。

// 例如，"ACGAATTCCG" 是一个 DNA序列 。
// 在研究 DNA 时，识别 DNA 中的重复序列非常有用。

// 给定一个表示DNA序列的字符串 s ，返回所有在DNA分子中出现不止一次的长度为10的序列(子字符串)。你可以按 任意顺序 返回答案。

// 不止出现一次长度为10的序列

func TestLeetCode187() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
func findRepeatedDnaSequences(s string) []string {
	dict := make(map[string]int)
	strs := []string{}
	for start := 0; start+10 <= len(s); start++ {
		str := s[start : start+10]
		if dict[str] == 1 {
			strs = append(strs, str)
		}
		dict[str]++
	}
	return strs
}
