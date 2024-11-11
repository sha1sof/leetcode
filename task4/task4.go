package task4

import (
	"fmt"
)

// A very poor result in memory and speed. Below are the best options in their
// category.
func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	resultMap := map[int]int{}
	for _, num := range nums1 {
		if _, ok := resultMap[num]; !ok {
			resultMap[num]++
		}
	}

	fmt.Println(resultMap)
	for _, num := range nums2 {
		_, ok := resultMap[num]
		if ok && resultMap[num] == 1 {
			resultMap[num]++
		}
	}

	fmt.Println(resultMap)
	for k, _ := range resultMap {
		if resultMap[k] > 1 {
			result = append(result, k)
		}
	}
	return result
}

// By speed.
/*func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	map2 := make(map[int]bool) //check lap lai
	res := make([]int, 0)
	for _, n := range nums1 {
		map1[n]++
	}

	for _, n := range nums2 {
		if map1[n] > 0 && !map2[n] {
			res = append(res, n)
			map2[n] = true
		}
	}
	return res
}*/

// By memory.
/*func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}

	for i := range nums1 {
		if hasValue(nums2, nums1[i]) && !hasValue(result, nums1[i]) {
			result = append(result, nums1[i])
		}
	}

	return result
}

func hasValue(nums []int, value int) bool {
	for i := range nums {
		if nums[i] == value {
			return true
		}
	}

	return false
}*/
