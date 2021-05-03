package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := isValid(tt.args.s); got != tt.want {
			t.Errorf("%v. isValid() = %v, want %v", tt.args.s, got, tt.want)
		}
	}
}
