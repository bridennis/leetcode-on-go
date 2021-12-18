package easy_valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "There is not valid pattern #1",
			args: struct{ s string }{s: "(){}}{"},
			want: false,
		},
		{
			name: "There is valid pattern #2",
			args: struct{ s string }{s: "()[]{}"},
			want: true,
		},
		{
			name: "There is not valid pattern #3",
			args: struct{ s string }{s: "(]"},
			want: false,
		},
		{
			name: "There is not valid pattern #4",
			args: struct{ s string }{s: "([)]"},
			want: false,
		},
		{
			name: "There is valid pattern #5",
			args: struct{ s string }{s: "{[]}"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
