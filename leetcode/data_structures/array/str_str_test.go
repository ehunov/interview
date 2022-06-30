package array

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				"hello",
				"ll",
			},
			2,
		},
		{
			args{
				"aaaaa",
				"bba",
			},
			-1,
		},
		{
			args{
				"a",
				"a",
			},
			0,
		},
		{
			args{
				"a",
				"",
			},
			0,
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
