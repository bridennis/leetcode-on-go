package easy_single_number

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Number #1", args: struct{ nums []int }{nums: []int{2, 2, 1}}, want: 1},
		{name: "Number #2", args: struct{ nums []int }{nums: []int{4, 1, 2, 1, 2}}, want: 4},
		{name: "Number #3", args: struct{ nums []int }{nums: []int{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
