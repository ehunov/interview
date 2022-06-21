package array

import "testing"

func TestDominantIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				[]int{3, 6, 1, 0},
			},
			1,
		},
		{
			"",
			args{
				[]int{1, 2, 3, 4},
			},
			-1,
		},
		{
			"",
			args{
				[]int{1},
			},
			0,
		},

		{
			"",
			args{
				[]int{},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
