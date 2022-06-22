package array

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{
				[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			args{
				[][]int{{1, 2}, {3, 4}},
			},
			[]int{1, 2, 3, 4},
		},
		{
			args{
				[][]int{{1}},
			},
			[]int{1},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
