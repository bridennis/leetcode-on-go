package medium_longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Longest substring #1",
			args: struct{ s string }{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "Longest substring #2",
			args: struct{ s string }{s: "bbbbb"},
			want: 1,
		},
		{
			name: "Longest substring #3",
			args: struct{ s string }{s: "pwwkew"},
			want: 3,
		},
		{
			name: "Longest substring #4",
			args: struct{ s string }{s: "ckilbkd"},
			want: 5,
		},
		{
			name: "Longest substring #5",
			args: struct{ s string }{s: "aab"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
