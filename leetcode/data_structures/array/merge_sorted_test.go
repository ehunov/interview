package array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		}, {
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		}, {
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{2, 0},
			1,
			[]int{1},
			1,
			[]int{1, 2},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := merge(tt.nums1, tt.m, tt.nums2, tt.n); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("merge() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
