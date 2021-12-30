package medium_combinations

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Combine #1",
			args: struct {
				n int
				k int
			}{n: 4, k: 2},
			want: [][]int{
				{4, 3},
				{4, 2},
				{3, 2},
				{4, 1},
				{3, 1},
				{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
