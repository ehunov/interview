package array

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	maxPairMinSum := 0
	numsLen := len(nums)
	for i := 0; i < numsLen; i += 2 {
		maxPairMinSum += nums[i]
	}

	return maxPairMinSum
}
