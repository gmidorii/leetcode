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
		// TODO: Add test cases.
		{
			args: args{s: "42"},
			want: 42,
		},
	}
	for _, tt := range tests {
		if got := myAtoi(tt.args.s); got != tt.want {
			t.Errorf("%q. myAtoi() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
