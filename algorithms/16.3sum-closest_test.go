package main

import "testing"

func Test_threeSumClosest(t *testing.T) {
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
			name: "",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 1, 1, 0},
				target: -100,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
			t.Errorf("%q. threeSumClosest() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
