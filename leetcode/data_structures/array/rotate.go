// https://leetcode.com/explore/learn/card/array-and-string/204/conclusion/1182/

package array

func rotate(nums []int, k int) []int {
	lenNums, curIndex, curNum, startIndex := len(nums), 0, nums[0], 0

	for i := 0; i < lenNums; i++ {
		newIndex := computeNewIndex(curIndex, lenNums, k)

		curNum, nums[newIndex] = nums[newIndex], curNum

		if newIndex == startIndex && newIndex < lenNums-1 {
			curIndex = newIndex + 1
			curNum = nums[curIndex]
			startIndex = curIndex
		} else {
			curIndex = newIndex
		}
	}

	return nums
}

// O(n)
func rotateWithAdditionalMemory(nums []int, k int) []int {
	lenNums := len(nums)
	result := make([]int, lenNums)
	for i := 0; i < lenNums; i++ {
		result[computeNewIndex(i, lenNums, k)] = nums[i]
	}

	return result
}

func computeNewIndex(curIndex int, lenNums int, k int) int {
	newIndex := curIndex + k
	for newIndex >= lenNums {
		newIndex -= lenNums
	}

	return newIndex
}

// O(n*k)
func rotateStepByStep(nums []int, k int) []int {
	lastNumsIndex := len(nums) - 1
	for i := 0; i < k; i++ {
		rotated := nums[lastNumsIndex]
		for j := 0; j < lastNumsIndex; j++ {
			nums[lastNumsIndex-j] = nums[lastNumsIndex-j-1]
		}
		nums[0] = rotated
	}

	return nums
}
