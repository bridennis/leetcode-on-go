package easy_reverse_string

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Reverse #1",
			args: struct{ s []byte }{s: []byte("hello")},
			want: []byte("olleh"),
		},
		{
			name: "Reverse #2",
			args: struct{ s []byte }{s: []byte("Hannah")},
			want: []byte("hannaH"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			got := tt.args.s
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
