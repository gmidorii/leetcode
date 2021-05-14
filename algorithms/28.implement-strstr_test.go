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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
			t.Errorf("%q. strStr() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
