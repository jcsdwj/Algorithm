/**
  @author: 伍萬
  @since: 2024/6/10
  @desc: //TODO
**/

package leetCode

import (
	"fmt"
	"sort"
)

// 881. 救生艇
// 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit
// 返回 承载所有人所需的最小船数

func LeetCode881() {
	people1 := []int{1, 2}
	people2 := []int{3, 2, 2, 1}
	people3 := []int{3, 5, 3, 4}
	fmt.Println(numRescueBoats(people1, 3))
	fmt.Println(numRescueBoats(people2, 3))
	fmt.Println(numRescueBoats(people3, 5))
}

func numRescueBoats(people []int, limit int) int {
	// 最小要尽可能和最大的匹配
	sort.Ints(people)
	start := 0
	end := len(people) - 1

	if people[start] > limit {
		return 0
	}

	count := 0
	for start <= end {
		// 所有都不满足
		if people[start] > limit {
			return count
		}
		// 起始节点就满足
		if people[start] == limit {
			count++
			start++
			continue
		}
		if people[start]+people[end] <= limit {
			count++
			start++
			end--
		} else {
			// 太大情况
			if people[end] <= limit {
				count++
			}

			// 尾指针往前移
			end--
		}

	}

	return count
}
