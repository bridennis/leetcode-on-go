package easy_number_of_1_bits

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Weight #1", args: struct{ num uint32 }{num: uint32(11)}, want: 3},
		{name: "Weight #2", args: struct{ num uint32 }{num: uint32(128)}, want: 1},
		{name: "Weight #2", args: struct{ num uint32 }{num: uint32(4294967293)}, want: 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
