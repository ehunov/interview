package array

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args{
				5,
			},
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			args{
				1,
			},
			[][]int{{1}},
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := generate(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
