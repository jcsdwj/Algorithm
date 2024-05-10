/**
  @author: 伍萬
  @since: 2024/5/10
  @desc: //TODO
**/

package nowcoder

func findTargetSumWays(nums []int, target int) int {

	// write code here

	if len(nums) == 0 {

		return 0

	}

	if len(nums) == 1 {

		if nums[0] == 0 && nums[0] == target {

			return 2

		}

		if nums[0] == target || nums[0] == -target {

			return 1

		}

		return 0

	}

	v := nums[0]

	nums = nums[1:]

	// 首项为正数情况

	zhengAnswer := findTargetSumWays(nums, target-v)

	// 首项为负数情况

	fuAnswer := findTargetSumWays(nums, target+v)

	return zhengAnswer + fuAnswer

} // pass
