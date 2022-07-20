package array

import "testing"

func TestPivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Integer pivot",
			args{
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			3,
		},
		{
			"No pivot",
			args{
				nums: []int{1, 2, 3},
			},
			-1,
		},
		{
			"First one is pivot",
			args{
				nums: []int{2, 1, -1},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
