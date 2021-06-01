package main

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			args: args{
				n: 4,
			},
			want: "1211",
		},
	}
	for _, tt := range tests {
		if got := countAndSay(tt.args.n); got != tt.want {
			t.Errorf("%v. countAndSay() = %v, want %v", tt.args.n, got, tt.want)
		}
	}
}
