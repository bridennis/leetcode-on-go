package easy_longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "There is common prefix",
			args: struct{ strs []string }{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "There is no common prefix",
			args: struct{ strs []string }{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "There is no data strings",
			args: struct{ strs []string }{strs: []string{}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
