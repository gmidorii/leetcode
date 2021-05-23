package main

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: []int{1, 5, 1},
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				nums: []int{2, 6, 5, 4, 3},
			},
			want: []int{3, 2, 4, 5, 6},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2},
			},
			want: []int{2, 1},
		},
		{
			name: "",
			args: args{
				nums: []int{5, 4, 7, 5, 3, 2},
			},
			want: []int{5, 5, 2, 3, 4, 7},
		},
		{
			name: "",
			args: args{
				nums: []int{2, 3, 1},
			},
			want: []int{3, 1, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 5, 1},
			},
			want: []int{5, 1, 1},
		},
		{
			name: "",
			args: args{
				nums: []int{2, 2, 7, 5, 4, 3, 2, 2, 1},
			},
			want: []int{2, 3, 1, 2, 2, 2, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		nextPermutation(tt.args.nums)
		if !reflect.DeepEqual(tt.want, tt.args.nums) {
			t.Errorf("want: %v, got: %v", tt.want, tt.args.nums)
		}
	}
}
