package leetCode

import (
	"fmt"
	"sort"
)

// 有点难 值得继续看

// 2071. 你可以安排的最多任务数目

// 给你 n 个任务和 m 个工人。每个任务需要一定的力量值才能完成，需要的力量值保存在下标从 0 开始的整数数组 tasks 中，
// 第 i 个任务需要 tasks[i] 的力量才能完成。每个工人的力量值保存在下标从 0 开始的整数数组 workers 中，第 j 个工人的力量值为 workers[j] 。
// 每个工人只能完成 一个 任务，且力量值需要 大于等于 该任务的力量要求值（即 workers[j] >= tasks[i] ）。

// 除此以外，你还有 pills 个神奇药丸，可以给 一个工人的力量值 增加 strength 。你可以决定给哪些工人使用药丸，但每个工人 最多 只能使用 一片 药丸。

// 给你下标从 0 开始的整数数组tasks 和 workers 以及两个整数 pills 和 strength ，请你返回最多有多少个任务可以被完成。

// 假设能完成k个任务，选最强的k个人去做最容易的k个任务
// 如果第1个人能完成任务k，则这k个任务都能完成

func TestLeetCode2071() {
	task1 := []int{3, 2, 1}
	worker1 := []int{0, 3, 3}
	fmt.Println(maxTaskAssign2(task1, worker1, 1, 1))
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	m, n := len(tasks), len(workers)

	return sort.Search(min(m, n)+1, func(k int) bool {
		if k > min(m, n) {
			return false
		}

		todo, p := []int{}, pills
		idx := 0 // 任务

		// 选最强的k个人
		for i := n - k; i < n; i++ {
			// 吃药可做任务加入列表
			for ; idx < k && workers[i]+strength >= tasks[idx]; idx++ {
				todo = append(todo, tasks[idx])
			}
			if len(todo) == 0 { // 没有任何任务
				return true
			}
			if workers[i] >= todo[0] { // 不吃药则做最小task
				todo = todo[1:]
			} else { // 吃药则做可做列表最大task
				if p == 0 {
					return true
				}
				todo = todo[:len(todo)-1]
				p--
			}
		}
		return false
	}) - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxTaskAssign2(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	Min := 0
	Max := min(len(tasks), len(workers))
	for Min < Max {
		mid := (Min + Max + 1) >> 1
		if check(mid, tasks, workers, pills, strength) {
			// 至少能完成mid个任务
			Min = mid
		} else {
			Max = mid - 1
		}
	}
	return Min
}

func check(mid int, tasks []int, workers []int, pills int, strength int) bool {
	index := len(workers) - 1
	queue := []int{}
	//工作力量递减遍历
	for i := mid - 1; i >= 0; i-- {
		// 先遍历大的任务
		task := tasks[i]
		// 假设每个工人都可以吃上药丸，将吃了药丸工人力量大于工作从大到小保存到队列中，只保留mid个
		for index >= 0 && workers[index]+strength >= task {
			// 放在末尾（大的先放）
			queue = append(queue, workers[index])
			index--
		}

		// 吃了药丸也不能完成工作
		if len(queue) == 0 {
			return false
		}

		//队伍最大力量不吃药丸可以完成工作
		if queue[0] >= task {
			queue = queue[1:]
		} else if pills-1 >= 0 {
			// 减掉一个药丸 让最小的吃
			queue = queue[:len(queue)-1]
			pills--
		} else {
			pills--
			return false
		}
	}
	return true
}
