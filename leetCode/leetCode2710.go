/**
  @author: 伍萬
  @since: 2024/6/29
  @desc: //TODO
**/

package leetCode

// 2710. 移除字符串中的尾随零

func removeTrailingZeros(num string) string {
	if num[len(num)-1] == '0' {
		return removeTrailingZeros(num[:len(num)-1])
	}
	return num
}
