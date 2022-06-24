package array

import "testing"

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{
				"11",
				"11",
			},
			"110",
		},
		{
			args{
				"11",
				"1",
			},
			"100",
		},
		{
			args{
				"11011",
				"10",
			},
			"11101",
		},
		{
			args{
				"10",
				"1011",
			},
			"1101",
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
