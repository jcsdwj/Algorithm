/**
  @author: 伍萬
  @since: 2024/3/5
  @desc: //TODO
**/

package nowcoder

func partition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sum := 0
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	if maxNum > sum/2 {
		return false
	}

	if sum%2 != 0 {
		return false
	}

	// dp[i][j]表示0-i位置元素是否存在和为j的元素
	// dp[i][j] = dp[i][j]dp[i-1][j-nums[i]] j>nums[i]
	// dp[i][j] = dp[i-1][j] j<nums[i]
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, sum/2+1)
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		for j := 1; j < sum%2; j++ {
			if j >= num {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][sum/2]
}
