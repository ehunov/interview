package tasks

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		firstStart  int
		firstEnd    int
		secondStart int
		secondEnd   int
		expected    []int
	}{
		{
			10,
			20,
			15,
			30,
			[]int{15, 20},
		},
		{
			15,
			30,
			10,
			20,
			[]int{15, 20},
		},
		{
			10,
			30,
			15,
			20,
			[]int{15, 20},
		},
		{
			15,
			20,
			10,
			30,
			[]int{15, 20},
		},
		{
			10,
			20,
			25,
			30,
			[]int{0, 0},
		},
		{
			25,
			30,
			10,
			20,
			[]int{0, 0},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := intersect(tt.firstStart, tt.firstEnd, tt.secondStart, tt.secondEnd); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("intersect() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
