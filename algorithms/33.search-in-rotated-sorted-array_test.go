package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		if got := search(tt.args.nums, tt.args.target); got != tt.want {
			t.Errorf("%q. search() = %v, want %v", tt.name, got, tt.want)
		}
		if got := search33(tt.args.nums, tt.args.target); got != tt.want {
			t.Errorf("%q. search() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
