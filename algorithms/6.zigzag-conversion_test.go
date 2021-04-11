package main

import (
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			args: args{s: "PAYPALISHIRING", numRows: 4},
			want: "PINALSIGYAHRPI",
		},
		{
			args: args{s: "A", numRows: 1},
			want: "A",
		},
	}
	for _, tt := range tests {
		if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
			t.Errorf("%q. convert() = %v, want %v", tt.args.s, got, tt.want)
		}
	}
}
