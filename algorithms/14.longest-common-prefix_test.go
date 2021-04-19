package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			args: args{strs: []string{"a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		if got := longestCommonPrefix(tt.args.strs); got != tt.want {
			t.Errorf("%v. longestCommonPrefix() = %v, want %v", tt.args.strs, got, tt.want)
		}
	}
}
