/**
  @author: 伍萬
  @since: 2024/2/24
  @desc: //TODO
**/

package nowcoder

import (
	"fmt"
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */

func TestNC83() {
	nums1 := []int{3, 2, -1, 4}
	nums2 := []int{-3, 0, -2}
	fmt.Println(maxProduct(nums1))
	fmt.Println(maxProduct(nums2))
}

func maxProduct(nums []int) int {

	if len(nums) < 2 {
		return nums[0]
	}

	// write code here
	// 以i结尾的最大最小值
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))

	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	maxValue := nums[0]

	for i := 1; i < len(nums); i++ {
		dpMax[i] = int(math.Max(math.Max(float64(nums[i]), float64(nums[i]*dpMax[i-1])), float64(nums[i]*dpMin[i-1])))
		dpMin[i] = int(math.Min(math.Min(float64(nums[i]), float64(nums[i]*dpMax[i-1])), float64(nums[i]*dpMin[i-1])))
		if dpMax[i] > maxValue {
			maxValue = dpMax[i]
		}
	}

	return maxValue
}
