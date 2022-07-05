// https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1301/

package array

func findMaxConsecutiveOnes(nums []int) int {
	consecutiveOnes, maxConsecutiveOnes := 0, 0

	for _, num := range nums {
		if num == 1 {
			consecutiveOnes++
		} else {
			consecutiveOnes = 0
		}

		if maxConsecutiveOnes < consecutiveOnes {
			maxConsecutiveOnes = consecutiveOnes
		}
	}

	return maxConsecutiveOnes
}
