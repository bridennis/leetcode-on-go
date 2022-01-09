package medium_unique_paths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Paths #1",
			args: struct {
				m int
				n int
			}{m: 3, n: 7},
			want: 28,
		},
		{
			name: "Paths #2",
			args: struct {
				m int
				n int
			}{m: 3, n: 2},
			want: 3,
		},
		{
			name: "Paths #3",
			args: struct {
				m int
				n int
			}{m: 1, n: 1},
			want: 1,
		},
		{
			name: "Paths #4",
			args: struct {
				m int
				n int
			}{m: 1, n: 10},
			want: 1,
		},
		{
			name: "Paths #5",
			args: struct {
				m int
				n int
			}{m: 10, n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
