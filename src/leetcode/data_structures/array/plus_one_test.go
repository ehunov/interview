package array

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{
				[]int{9, 9},
			},
			[]int{1, 0, 0},
		},
		{
			args{
				[]int{1, 2, 3},
			},
			[]int{1, 2, 4},
		},
		{
			args{
				[]int{4, 3, 2, 1},
			},
			[]int{4, 3, 2, 2},
		},
		{
			args{
				[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
			},
			[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
