package array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{
				[]int{2, 7, 11, 15},
				9,
			},
			[]int{1, 2},
		}, {
			args{
				[]int{2, 3, 4},
				6,
			},
			[]int{1, 3},
		}, {
			args{
				[]int{-1, 0},
				-1,
			},
			[]int{1, 2},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
