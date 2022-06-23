package array

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{
				[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			args{
				[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
