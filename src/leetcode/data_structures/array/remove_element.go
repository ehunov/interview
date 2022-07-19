// https://leetcode.com/explore/learn/card/array-and-string/205/array-two-pointer-technique/1151/

package array

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			continue
		}

		i++
	}

	return i
}
