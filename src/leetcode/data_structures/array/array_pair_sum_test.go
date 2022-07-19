package array

import "testing"

func TestArrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				[]int{1, 4, 3, 2},
			},
			4,
		},
		{
			args{
				[]int{6, 2, 6, 5, 1, 2},
			},
			9,
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
