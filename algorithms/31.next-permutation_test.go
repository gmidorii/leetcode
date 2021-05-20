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
	}
	for _, tt := range tests {
		nextPermutation(tt.args.nums)
		if !reflect.DeepEqual(tt.want, tt.args.nums) {
			t.Errorf("want: %v, got: %v", tt.want, tt.args.nums)
		}
	}
}
