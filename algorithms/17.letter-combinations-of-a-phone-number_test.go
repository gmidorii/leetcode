package main

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "",
			args: args{digits: ""},
			want: []string{},
		},
		{
			name: "",
			args: args{digits: "2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. letterCombinations() = %v, want %v", tt.name, got, tt.want)
		}
		if got := letterCombinations2(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. letterCombinations() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
