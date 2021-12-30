package easy_letter_case_permutation

import (
	"reflect"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Letter case permutation #1",
			args: struct{ s string }{s: "a1b2"},
			want: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			name: "Letter case permutation #2",
			args: struct{ s string }{s: "3z4"},
			want: []string{"3z4", "3Z4"},
		},
		{
			name: "Letter case permutation #3",
			args: struct{ s string }{s: "abc"},
			want: []string{"abc", "abC", "aBc", "aBC", "Abc", "AbC", "ABc", "ABC"},
		},
		{
			name: "Letter case permutation #4",
			args: struct{ s string }{s: "1a3"},
			want: []string{"1a3", "1A3"},
		},
		{
			name: "Letter case permutation #5",
			args: struct{ s string }{s: ""},
			want: []string{""},
		},
		{
			name: "Letter case permutation #6",
			args: struct{ s string }{s: "123"},
			want: []string{"123"},
		},
		{
			name: "Letter case permutation #7",
			args: struct{ s string }{s: "C"},
			want: []string{"c", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
