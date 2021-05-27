package main

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. searchRange() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
