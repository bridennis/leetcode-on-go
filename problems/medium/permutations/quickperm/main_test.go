package medium_quickperm_permutations

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Permute #1",
			args: struct{ nums []int }{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}},
		},
		{
			name: "Permute #2",
			args: struct{ nums []int }{nums: []int{0, 1}},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "Permute #1",
			args: struct{ nums []int }{nums: []int{1}},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
