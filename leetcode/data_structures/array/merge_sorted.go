// https://leetcode.com/problems/merge-sorted-array/

package array

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i := m + n - 1
	m--
	n--

	for n >= 0 {
		if m >= 0 && nums1[m] >= nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}

	return nums1
}
