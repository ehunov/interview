// https://leetcode.com/explore/learn/card/array-and-string/201/introduction-to-array/1147/

package array

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := 0
	maxIndex := 0
	closestToMax := 0

	for i, num := range nums {
		if num >= max {
			closestToMax = max
			max = num
			maxIndex = i
		} else if num >= closestToMax {
			closestToMax = num
		}
	}

	if closestToMax*2 > max {
		return -1
	}

	return maxIndex
}
