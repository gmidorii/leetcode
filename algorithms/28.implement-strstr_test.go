package main

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
			name: "",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				haystack: "",
				needle:   "a",
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
			t.Errorf("%q. strStr() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
