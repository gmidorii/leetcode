package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		nums:   []int{1, 3, 5, 6},
		// 		target: 5,
		// 	},
		// 	want: 2,
		// },
		// {
		// 	args: args{
		// 		nums:   []int{1, 3, 5, 6},
		// 		target: 2,
		// 	},
		// 	want: 1,
		// },
		// {
		// 	args: args{
		// 		nums:   []int{1, 3, 5, 6},
		// 		target: 7,
		// 	},
		// 	want: 4,
		// },
		// {
		// 	args: args{
		// 		nums:   []int{1, 3, 5, 6},
		// 		target: 0,
		// 	},
		// 	want: 0,
		// },
		{
			args: args{
				nums:   []int{1},
				target: 2,
			},
			want: 1,
		},
		{
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := searchInsert2(tt.args.nums, tt.args.target); got != tt.want {
			t.Errorf("%q. searchInsert() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
