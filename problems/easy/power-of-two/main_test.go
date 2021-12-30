package easy_power_of_two

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Pow #1",
			args: struct{ n int }{n: 1},
			want: true,
		},
		{
			name: "Pow #2",
			args: struct{ n int }{n: 16},
			want: true,
		},
		{
			name: "Pow #3",
			args: struct{ n int }{n: 3},
			want: false,
		},
		{
			name: "Pow #4",
			args: struct{ n int }{n: 6},
			want: false,
		},
		{
			name: "Pow #5",
			args: struct{ n int }{n: 12},
			want: false,
		},
		{
			name: "Pow #6",
			args: struct{ n int }{n: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
