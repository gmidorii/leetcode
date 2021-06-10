package main

import (
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}
	for _, tt := range tests {
		got := combinationSum2(tt.args.candidates, tt.args.target)
		if !matchMatrix(got, tt.want) {
			t.Errorf("%q. combinationSum2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func matchMatrix(ss, targets [][]int) bool {
	if len(ss) != len(targets) {
		return false
	}
	matchCount := 0
	for _, s := range ss {
		for _, t := range targets {
			if matchSlice2(s, t) {
				matchCount++
			}
		}
	}
	if matchCount == len(targets) {
		return true
	}
	return false
}
