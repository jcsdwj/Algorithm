/**
  @author: 伍萬
  @since: 2024/7/16
  @desc: //TODO
**/

package leetCode

// 2956. 找到两个数组中的公共元素

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	dictNums1 := make(map[int]bool)
	dictNums2 := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		dictNums1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		dictNums2[nums2[i]] = true
	}

	count1 := 0
	count2 := 0
	for _, v := range nums1 {
		if dictNums2[v] {
			count1++
		}
	}
	for _, v := range nums2 {
		if dictNums1[v] {
			count2++
		}
	}
	return []int{count1, count2}
}
