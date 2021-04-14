package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{x: 121},
			want: true,
		},
		{
			args: args{x: 122},
			want: false,
		},
		{
			args: args{x: -121},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := isPalindrome(tt.args.x); got != tt.want {
			t.Errorf("%v. isPalindrome() = %v, want %v", tt.args.x, got, tt.want)
		}
	}
}
