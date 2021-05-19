package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			args: args{
				dividend: 0,
				divisor:  1,
			},
			want: 0,
		},
		{
			args: args{
				dividend: 1,
				divisor:  1,
			},
			want: 1,
		},
		{
			args: args{
				dividend: 7,
				divisor:  -10,
			},
			want: 0,
		},
		{
			args: args{
				dividend: -1,
				divisor:  1,
			},
			want: -1,
		},
		{
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: 2147483647,
		},
		{
			args: args{
				dividend: -2147483648,
				divisor:  1,
			},
			want: -2147483648,
		},
		{
			args: args{
				dividend: -2147483648,
				divisor:  -3,
			},
			want: 715827882,
		},
	}
	for _, tt := range tests {
		if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
			t.Errorf("%v. divide() = %v, want %v", tt.args.dividend, got, tt.want)
		}
		if got := divide1(tt.args.dividend, tt.args.divisor); got != tt.want {
			t.Errorf("%v/%v. divide1() = %v, want %v", tt.args.dividend, tt.args.divisor, got, tt.want)
		}
	}
}
