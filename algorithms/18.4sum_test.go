package main

import (
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
		{
			name: "",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			want: [][]int{[]int{2, 2, 2, 2}},
		},
		{
			name: "",
			args: args{
				nums:   []int{-2, -1, -1, 1, 1, 2, 2},
				target: 0,
			},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-1, -1, 1, 1}},
		},
	}
	for _, tt := range tests {
		if got := fourSum(tt.args.nums, tt.args.target); !matchSliceOnSlice(tt.want, got) {
			t.Errorf("%q. fourSum() = %v, want %v", tt.name, got, tt.want)
		}
		if got := fourSum2(tt.args.nums, tt.args.target); !matchSliceOnSlice(tt.want, got) {
			t.Errorf("%q. fourSum2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_distinct(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				nums: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
			},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
		{
			name: "duplicate",
			args: args{
				nums: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}, []int{-1, 0, 0, 1}},
			},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
		{
			name: "duplicate",
			args: args{
				nums: [][]int{[]int{-1, 0, 1, 0}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
			},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}},
		},
		{
			name: "duplicate",
			args: args{
				nums: [][]int{[]int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
			},
			want: [][]int{[]int{2, 2, 2, 2}},
		},
		{
			name: "duplicate",
			args: args{
				nums: [][]int{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}, []int{-1, -1, 1, 1}},
			},
			want: [][]int{[]int{-2, -1, 1, 2}, []int{-1, -1, 1, 1}},
		},
	}
	for _, tt := range tests {
		if got := distinct(tt.args.nums); !matchSliceOnSlice(tt.want, got) {
			t.Errorf("%q. distinct() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
