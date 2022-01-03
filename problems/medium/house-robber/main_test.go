package medium_house_robber

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Rob #1", args: struct{ nums []int }{nums: []int{1, 2, 3, 1}}, want: 4},
		{name: "Rob #2", args: struct{ nums []int }{nums: []int{2, 7, 9, 3, 1}}, want: 12},
		{name: "Rob #3", args: struct{ nums []int }{nums: []int{100, 1, 1, 100}}, want: 200},
		{name: "Rob #4", args: struct{ nums []int }{nums: []int{1, 100, 100}}, want: 101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
