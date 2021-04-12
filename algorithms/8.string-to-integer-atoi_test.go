package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{s: "42"},
			want: 42,
		},
		{
			args: args{s: "   -42"},
			want: -42,
		},
		{
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			args: args{s: "+-12"},
			want: 0,
		},
		{
			args: args{s: "9223372036854775808"},
			want: 2147483647,
		},
		{
			args: args{s: "-9223372036854775808"},
			want: -2147483648,
		},
		{
			args: args{s: "18446744073709551617"},
			want: 2147483647,
		},
		{
			args: args{s: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"},
			want: 2147483647,
		},
		{
			args: args{s: "-000000000000000000000000000000000000000000000000001"},
			want: -1,
		},
	}
	for _, tt := range tests {
		if got := myAtoi(tt.args.s); got != tt.want {
			t.Errorf("%q. myAtoi() = %v, want %v", tt.args.s, got, tt.want)
		}
	}
}