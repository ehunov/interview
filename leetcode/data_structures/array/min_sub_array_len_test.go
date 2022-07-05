package array

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			[]int{2, 3, 1, 2, 4, 3},
			7,
			2,
		},
		{
			[]int{1, 4, 4},
			4,
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			11,
			0,
		},
		{
			[]int{1, 2, 3, 4, 5},
			15,
			5,
		},
		{
			[]int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			213,
			8,
		},
		{
			[]int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			15,
			2,
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.expected {
				t.Errorf("minSubArrayLen() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
