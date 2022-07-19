package array

import (
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		args         args
		want         int
		expectedNums []int
	}{
		{
			args{
				[]int{3, 2, 2, 3},
				3,
			},
			2,
			[]int{2, 2},
		},
		{
			args{
				[]int{0, 1, 2, 2, 3, 0, 4, 2},
				2,
			},
			5,
			[]int{0, 1, 4, 0, 3},
		},
		{
			args{
				[]int{2},
				3,
			},
			1,
			[]int{2},
		},
		{
			args{
				[]int{},
				0,
			},
			0,
			[]int{},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			result := removeElement(tt.args.nums, tt.args.val)

			assert(t, result, tt.want)
			sort.Slice(tt.args.nums[0:result], func(i, j int) bool {
				return tt.args.nums[i] < tt.args.nums[j]
			})
			sort.Slice(tt.expectedNums, func(i, j int) bool {
				return tt.expectedNums[i] < tt.expectedNums[j]
			})

			for i := 0; i < result; i++ {
				assert(t, tt.args.nums[i], tt.expectedNums[i])
			}
		})
	}
}

func assert(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}
