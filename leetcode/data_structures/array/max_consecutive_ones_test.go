package array

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1, 1, 0, 1, 1, 1},
			3,
		},
		{
			[]int{1, 0, 1, 1, 0, 1},
			2,
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.nums); got != tt.expected {
				t.Errorf("findMaxConsecutiveOnes() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
