package easy_reverse_bits

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Reverse #1",
			args: struct{ num uint32 }{num: 43261596},
			want: 964176192,
		},
		{
			name: "Reverse #2",
			args: struct{ num uint32 }{num: 4294967293},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
