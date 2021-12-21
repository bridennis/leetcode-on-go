package easy_reverse_words_in_a_string_iii

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Reverse words #1",
			args: struct{ s string }{s: "Let's take LeetCode contest"},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "Reverse words #2",
			args: struct{ s string }{s: "God Ding"},
			want: "doG gniD",
		},
		{
			name: "Reverse words #3",
			args: struct{ s string }{s: "a"},
			want: "a",
		},
		{
			name: "Reverse words #4",
			args: struct{ s string }{s: "abc"},
			want: "cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
