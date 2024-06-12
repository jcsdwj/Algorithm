/**
  @author: 伍萬
  @since: 2024/6/12
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"math"
)

// 2806. 取整购买后的账户余额

func LeetCode2806() {
	fmt.Println(accountBalanceAfterPurchase(15))
}

func accountBalanceAfterPurchase(purchaseAmount int) int {
	return 100 - int(math.Round(float64(purchaseAmount)/10))*10
}
