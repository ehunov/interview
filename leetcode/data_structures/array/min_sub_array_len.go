// https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1299/

package array

func minSubArrayLen(target int, nums []int) int {
	lenNums := len(nums)
	minLength, left, sum := lenNums+1, 0, 0

	for i, num := range nums {
		sum += num
		for ; sum >= target; left++ {
			curMin := i + 1 - left
			if curMin < minLength {
				minLength = curMin
			}
			sum -= nums[left]
		}
	}

	if minLength > lenNums {
		return 0
	}

	return minLength
}
