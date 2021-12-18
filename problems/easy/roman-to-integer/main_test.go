package easy_roman_to_integer

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should get the right result for LVIII",
			args: struct{ s string }{s: "LVIII"},
			want: 58,
		},
		{
			name: "Should get the right result for MCMXCIV",
			args: struct{ s string }{s: "MCMXCIV"},
			want: 1994,
		},
		{
			name: "Should get the right result for IV",
			args: struct{ s string }{s: "IV"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
