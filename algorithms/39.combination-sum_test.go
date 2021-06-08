package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
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
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
		{
			args: args{
				candidates: []int{1},
				target:     1,
			},
			want: [][]int{
				{1},
			},
		},
		{
			args: args{
				candidates: []int{1},
				target:     2,
			},
			want: [][]int{
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. combinationSum2() = %v, want %v", tt.name, got, tt.want)
		}
		if got := combinationSum3(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. combinationSum3() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
