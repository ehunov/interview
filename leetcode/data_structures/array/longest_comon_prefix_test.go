package array

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{
				[]string{"flower", "flow", "flight"},
			},
			"fl",
		},
		{
			args{
				[]string{"dog", "racecar", "car"},
			},
			"",
		},
		{
			args{
				[]string{"ab", "a"},
			},
			"a",
		},
	}
	for i, tt := range tests {
		t.Run(string(rune(i)), func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
