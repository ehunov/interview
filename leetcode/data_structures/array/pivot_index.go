// https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1144/

package array

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	leftSum := 0
	var rightSum int

	for _, num := range nums {
		rightSum += num
	}

	for i, num := range nums {
		rightSum -= num
		if rightSum == leftSum {
			return i
		}
		leftSum += num
	}

	return -1
}
