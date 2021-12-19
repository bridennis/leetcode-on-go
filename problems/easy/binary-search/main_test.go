package easy_binary_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Search #1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "Search #2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "Search #3",
			args: args{
				nums:   []int{1, 2},
				target: 2,
			},
			want: 1,
		},
		{
			name: "Search #4",
			args: args{
				nums:   []int{1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Search #5",
			args: args{
				nums:   []int{10},
				target: 10,
			},
			want: 0,
		},
		{
			name: "Search #6",
			args: args{
				nums:   []int{1, 2, 3},
				target: 3,
			},
			want: 2,
		},
		{
			name: "Search #7",
			args: args{
				nums:   []int{-1, 0, 5},
				target: -1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
