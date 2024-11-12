package task3

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	answer := make([]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(queries); i++ {
		sum := 0
		result := 0
		for j := 0; j < len(nums); j++ {
			if sum+nums[j] <= queries[i] && j != len(nums)-1 {
				sum += nums[j]
				result++
			} else if sum+nums[j] <= queries[i] && j == len(nums)-1 {
				result++
				answer = append(answer, result)
				break
			} else {
				answer = append(answer, result)
				break
			}
		}
	}
	return answer
}
