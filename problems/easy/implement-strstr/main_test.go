package easy_implement_strstr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "strStr #1",
			args: struct {
				haystack string
				needle   string
			}{haystack: "hello", needle: "ll"},
			want: 2,
		},
		{
			name: "strStr #2",
			args: struct {
				haystack string
				needle   string
			}{haystack: "aaaaa", needle: "bba"},
			want: -1,
		},
		{
			name: "strStr #3",
			args: struct {
				haystack string
				needle   string
			}{haystack: "", needle: ""},
			want: 0,
		},
		{
			name: "strStr #4",
			args: struct {
				haystack string
				needle   string
			}{haystack: "abc", needle: "c"},
			want: 2,
		},
		{
			name: "strStr #5",
			args: struct {
				haystack string
				needle   string
			}{haystack: "", needle: "a"},
			want: -1,
		},
		{
			name: "strStr #6",
			args: struct {
				haystack string
				needle   string
			}{haystack: "a", needle: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strStr("hello", "ll")
	}
}
