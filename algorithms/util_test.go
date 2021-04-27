package main

import "testing"

func matchSlice(src []int, target []int) bool {
	var tmp []int
	for _, s := range src {
		for _, t := range target {
			if s != t {
				tmp = append(tmp, t)
			}
		}
		target = tmp
		tmp = []int{}
	}
	return len(target) == 0
}

func Test_matchSlice(t *testing.T) {
	type args struct {
		src    []int
		target []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				src:    []int{1, 3, 2, 4},
				target: []int{1, 2, 3, 4},
			},
			want: true,
		},
		{
			args: args{
				src:    []int{1, 3, 2, 5},
				target: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := matchSlice(tt.args.src, tt.args.target); tt.want != got {
			t.Errorf("not match slice src=%v, target=%v, got=%v, want=%v", tt.args.src, tt.args.target, got, tt.want)
		}
	}
}

func matchSliceOnSlice(src [][]int, target [][]int) bool {
	if len(src) != len(target) {
		return false
	}
	var tmp [][]int
	for _, ss := range src {
		for _, ts := range target {
			if !matchSlice(ss, ts) {
				tmp = append(tmp, ts)
			}
		}
		target = tmp
		tmp = [][]int{}
	}
	return len(target) == 0
}

func Test_matchSliceOnSlice(t *testing.T) {
	type args struct {
		src    [][]int
		target [][]int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				src:    [][]int{[]int{1, 3, 2, 4}},
				target: [][]int{[]int{1, 2, 3, 4}},
			},
			want: true,
		},
		{
			args: args{
				src:    [][]int{[]int{1, 3, 2, 5}},
				target: [][]int{[]int{1, 2, 3, 4}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := matchSliceOnSlice(tt.args.src, tt.args.target); tt.want != got {
			t.Errorf("not match slice on slice src=%v, target=%v, got=%v, want=%v", tt.args.src, tt.args.target, got, tt.want)
		}
	}
}
