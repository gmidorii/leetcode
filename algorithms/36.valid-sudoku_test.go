package main

import "testing"

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				board: [][]byte{
					[]byte(string([]rune{'5', '3', '.', '.', '7', '.', '.', '.', '.'})),
					[]byte(string([]rune{'6', '.', '.', '1', '9', '5', '.', '.', '.'})),
					[]byte(string([]rune{'.', '9', '8', '.', '.', '.', '.', '6', '.'})),
					[]byte(string([]rune{'8', '.', '.', '.', '6', '.', '.', '.', '3'})),
					[]byte(string([]rune{'4', '.', '.', '8', '.', '3', '.', '.', '1'})),
					[]byte(string([]rune{'7', '.', '.', '.', '2', '.', '.', '.', '6'})),
					[]byte(string([]rune{'.', '6', '.', '.', '.', '.', '2', '8', '.'})),
					[]byte(string([]rune{'.', '.', '.', '4', '1', '9', '.', '.', '5'})),
					[]byte(string([]rune{'.', '.', '.', '.', '8', '.', '.', '7', '9'})),
				},
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				board: [][]byte{
					[]byte(string([]rune{'8', '3', '.', '.', '7', '.', '.', '.', '.'})),
					[]byte(string([]rune{'6', '.', '.', '1', '9', '5', '.', '.', '.'})),
					[]byte(string([]rune{'.', '9', '8', '.', '.', '.', '.', '6', '.'})),
					[]byte(string([]rune{'8', '.', '.', '.', '6', '.', '.', '.', '3'})),
					[]byte(string([]rune{'4', '.', '.', '8', '.', '3', '.', '.', '1'})),
					[]byte(string([]rune{'7', '.', '.', '.', '2', '.', '.', '.', '6'})),
					[]byte(string([]rune{'.', '6', '.', '.', '.', '.', '2', '8', '.'})),
					[]byte(string([]rune{'.', '.', '.', '4', '1', '9', '.', '.', '5'})),
					[]byte(string([]rune{'.', '.', '.', '.', '8', '.', '.', '7', '9'})),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := isValidSudoku(tt.args.board); got != tt.want {
			t.Errorf("%q. isValidSudoku() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
