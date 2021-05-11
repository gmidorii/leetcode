package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2},
			},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1},
			},
			want:     2,
			wantNums: []int{1},
		},
	}
	for _, tt := range tests {
		before := make([]int, len(tt.args.nums))
		n := copy(before, tt.args.nums)
		t.Log(before)
		t.Log(n)
		t.Log(tt.args.nums)
		if got := removeDuplicates(tt.args.nums); got != tt.want {
			t.Errorf("%q. removeDuplicates() = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(tt.args.nums, tt.wantNums) {
			t.Errorf("%q. removeDuplicates() = %v, want %v", tt.name, tt.args.nums, tt.wantNums)
		}
	}
}
