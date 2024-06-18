/**
  @author: 伍萬
  @since: 2024/6/18
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"strconv"
	"strings"
)

//2288. 价格减免

func TestLeetCode2288() {
	fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 50))
	fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100))
}

func discountPrices(sentence string, discount int) string {
	// 保存数字的起始和结束位置
	words := strings.Split(sentence, " ")
	for i, word := range words {
		// 是这个开头
		if strings.HasPrefix(word, "$") {
			price, err := strconv.Atoi(word[1:])
			if err == nil {
				discountPrice := float64(price) * (1 - float64(discount)/100)
				words[i] = fmt.Sprintf("$%.2f", discountPrice)
			}
		}
	}
	return strings.Join(words, " ")
}
