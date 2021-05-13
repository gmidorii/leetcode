package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
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
				nums: []int{3, 2, 2, 3},
				val:  2,
			},
			want:     2,
			wantNums: []int{3, 3, 2, 3},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:     5,
			wantNums: []int{0, 1, 3, 0, 4, 0, 4, 2},
		},
	}
	for _, tt := range tests {
		if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
			t.Errorf("%q. removeElement() = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(tt.args.nums, tt.wantNums) {
			t.Errorf("%q. removeElement() = %v, want %v", tt.name, tt.args.nums, tt.wantNums)
		}
	}
}
