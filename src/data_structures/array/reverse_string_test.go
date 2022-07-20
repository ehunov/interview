package array

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		args args
		want []byte
	}{
		{
			args{
				[]byte("hello"),
			},
			[]byte("olleh"),
		},
		{
			args{
				[]byte("Hannah"),
			},
			[]byte("hannaH"),
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := reverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
