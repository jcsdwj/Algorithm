/**
  @author: 伍萬
  @since: 2024/6/24
  @desc: //TODO
**/

package common

func quickSort(arr []int) {

}

// 方便的写法
func quickSortSolve(source []int) []int {
	if len(source) < 2 {
		return source
	}
	pivot := source[0]
	var left, right []int
	for _, value := range source[1:] {
		if value < pivot {
			left = append(left, value) // 比value小的放左边
		} else {
			right = append(right, value) // 比value小的放右边
		}
	}
	mid := append([]int{pivot}, quickSortSolve(right)...)
	return append(quickSortSolve(right), mid...)
}

func quickSortSolve2(nums []int) {

}
