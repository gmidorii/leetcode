package main

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{[]int{-1, 0, 1}, []int{-1, 2, -1}},
		},
		{
			args: args{nums: []int{}},
			want: [][]int{},
		},
		{
			args: args{nums: []int{0}},
			want: [][]int{},
		},
		{
			args: args{nums: []int{0, 0, 0}},
			want: [][]int{[]int{0, 0, 0}},
		},
		{
			args: args{nums: []int{0, 0, 0, 0}},
			want: [][]int{[]int{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. threeSum() = %v, want %v", tt.args.nums, got, tt.want)
		}
	}
}
