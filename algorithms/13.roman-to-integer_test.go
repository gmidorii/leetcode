package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{s: "III"},
			want: 3,
		},
		{
			args: args{s: "IV"},
			want: 4,
		},
		{
			args: args{s: "IX"},
			want: 9,
		},
		{
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
		{
			args: args{s: "MMMXLV"},
			want: 3045,
		},
	}
	for _, tt := range tests {
		if got := romanToInt(tt.args.s); got != tt.want {
			t.Errorf("%v. romanToInt() = %v, want %v", tt.args.s, got, tt.want)
		}
	}
}
