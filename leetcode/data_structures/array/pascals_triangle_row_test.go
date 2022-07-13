package array

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		expected []int
	}{
		{
			0,
			[]int{1},
		},
		{
			1,
			[]int{1, 1},
		},
		{
			2,
			[]int{1, 2, 1},
		},
		{
			3,
			[]int{1, 3, 3, 1},
		},
		{
			4,
			[]int{1, 4, 6, 4, 1},
		},
		{
			5,
			[]int{1, 5, 10, 10, 5, 1},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("getRow() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
